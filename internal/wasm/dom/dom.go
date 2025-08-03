// Package dom provides a Go-friendly wrapper
// for interacting with the JavaScript DOM API.
package dom

import (
	"encoding/json"
	"errors"
	"syscall/js"
)

type (
	// DOM represents a JavaScript DOM object.
	DOM struct {
		jsValue js.Value
	}
)

// Bool returns the DOM object as a boolean.
func (d *DOM) Bool() bool {
	return d.jsValue.Bool()
}

// String returns the DOM object as a string.
func (d *DOM) String() string {
	return d.jsValue.String()
}

// Int returns the DOM object as an integer.
func (d *DOM) Int() int {
	return d.jsValue.Int()
}

// Float returns the DOM object as a float.
func (d *DOM) Float() float64 {
	return d.jsValue.Float()
}

// Truthy returns the "truthiness" of the DOM object.
func (d *DOM) Truthy() bool {
	return d.jsValue.Truthy()
}

// Get retrieves a property from the DOM object.
func (d *DOM) Get(property string) *DOM {
	return &DOM{
		jsValue: d.jsValue.Get(property),
	}
}

// Set sets a property on the DOM object.
func (d *DOM) Set(property string, value any) {
	d.jsValue.Set(property, unwrapValue(value))
}

// Call invokes a method on the DOM object.
func (d *DOM) Call(method string, args ...any) *DOM {
	return &DOM{
		jsValue: d.jsValue.Call(method, unwrapValues(args)...),
	}
}

// Await awaits the resolution of the DOM object.
// It must be called on a Promise-returning method.
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

// GetGlobal returns the current global object.
func GetGlobal() *DOM {
	return &DOM{
		jsValue: js.Global(),
	}
}

// Fetch fetches a URL and unmarshals the response into a struct.
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

// unwrapValue unwraps a value into a JavaScript value.
func unwrapValue(value any) any {
	v, ok := value.(*DOM)
	if ok {
		return v.jsValue
	}
	return value
}

// unwrapValues unwraps a slice values into a slice of JavaScript values.
func unwrapValues(values []any) []any {
	unwrappedValues := make([]any, len(values))
	for i, v := range values {
		unwrappedValues[i] = unwrapValue(v)
	}
	return unwrappedValues
}

// onFulfilledCallback returns a JavaScript function that is called when a Promise is fulfilled.
func onFulfilledCallback(valueChan chan<- *DOM) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		valueChan <- &DOM{
			jsValue: args[0],
		}
		return nil
	})
}

// onRejectedCallback returns a JavaScript function that is called when a Promise is rejected.
func onRejectedCallback(errorChan chan<- error) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		errorChan <- errors.New(args[0].String())
		return nil
	})
}
