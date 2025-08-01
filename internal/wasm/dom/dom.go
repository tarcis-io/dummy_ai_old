package dom

import (
	"encoding/json"
	"errors"
	"syscall/js"
)

type (
	DOM struct {
		jsValue js.Value
	}
)

func (d *DOM) Bool() bool {
	return d.jsValue.Bool()
}

func (d *DOM) String() string {
	return d.jsValue.String()
}

func (d *DOM) Int() int {
	return d.jsValue.Int()
}

func (d *DOM) Float() float64 {
	return d.jsValue.Float()
}

func (d *DOM) Truthy() bool {
	return d.jsValue.Truthy()
}

func (d *DOM) Get(property string) *DOM {
	return &DOM{
		jsValue: d.jsValue.Get(property),
	}
}

func (d *DOM) Set(property string, value any) {
	d.jsValue.Set(property, unwrapValue(value))
}

func (d *DOM) Call(method string, args ...any) *DOM {
	return &DOM{
		jsValue: d.jsValue.Call(method, unwrapValues(args)...),
	}
}

func (d *DOM) Await() (*DOM, error) {
	valueChan := make(chan *DOM)
	errorChan := make(chan error)
	onFulfilled := onFulfilledCallback(valueChan)
	defer onFulfilled.Release()
	onRejected := onRejectedCallback(errorChan)
	defer onRejected.Release()
	d.Call("then", onFulfilled, onRejected)
	select {
	case v := <-valueChan:
		return v, nil
	case e := <-errorChan:
		return nil, e
	}
}

func GetGlobal() *DOM {
	return &DOM{
		jsValue: js.Global(),
	}
}

func Fetch[T any](url string) (*T, error) {
	v, err := GetGlobal().Call("fetch", url).Await()
	if err != nil {
		return nil, err
	}
	v, err = v.Call("text").Await()
	if err != nil {
		return nil, err
	}
	var t T
	err = json.Unmarshal([]byte(v.String()), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func unwrapValue(value any) any {
	v, ok := value.(*DOM)
	if ok {
		return v.jsValue
	}
	return value
}

func unwrapValues(values []any) []any {
	unwrappedValues := make([]any, len(values))
	for i, v := range values {
		unwrappedValues[i] = unwrapValue(v)
	}
	return unwrappedValues
}

func onFulfilledCallback(valueChan chan<- *DOM) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		valueChan <- &DOM{
			jsValue: args[0],
		}
		return nil
	})
}

func onRejectedCallback(errorChan chan<- error) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		errorChan <- errors.New(args[0].String())
		return nil
	})
}
