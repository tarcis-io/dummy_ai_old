package dom

import (
	"errors"
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func (d *DOM) Bool() bool {
	return d.jsObject.Bool()
}

func (d *DOM) String() string {
	return d.jsObject.String()
}

func (d *DOM) Int() int {
	return d.jsObject.Int()
}

func (d *DOM) Float() float64 {
	return d.jsObject.Float()
}

func (d *DOM) Truthy() bool {
	return d.jsObject.Truthy()
}

func (d *DOM) Get(property string) *DOM {
	return &DOM{
		jsObject: d.jsObject.Get(property),
	}
}

func (d *DOM) Set(property string, value any) {
	d.jsObject.Set(property, unwrapValue(value))
}

func (d *DOM) Call(method string, args ...any) *DOM {
	return &DOM{
		jsObject: d.jsObject.Call(method, unwrapValues(args)...),
	}
}

func (d *DOM) Await() (*DOM, error) {
	v := make(chan *DOM)
	defer close(v)
	onResolve := js.FuncOf(func(this js.Value, args []js.Value) any {
		v <- &DOM{
			jsObject: args[0],
		}
		return nil
	})
	err := make(chan error)
	defer close(err)
	onReject := js.FuncOf(func(this js.Value, args []js.Value) any {
		err <- errors.New(args[0].String())
		return nil
	})
	d.Call("then", onResolve, onReject)
	select {
	case v := <-v:
		return v, nil
	case err := <-err:
		return nil, err
	}
}

func GetGlobal() *DOM {
	return &DOM{
		jsObject: js.Global(),
	}
}

func unwrapValue(value any) any {
	v, ok := value.(*DOM)
	if ok {
		return v.jsObject
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
