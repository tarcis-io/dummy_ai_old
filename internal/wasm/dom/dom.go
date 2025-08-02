// Package dom provides a Go-friendly wrapper for interacting with the JavaScript DOM.
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

// Get retrieves a property from the DOM object and returns it as new *DOM object.
func (d *DOM) Get(property string) *DOM {
	return &DOM{
		jsValue: d.jsValue.Get(property),
	}
}

// Set sets a property of the DOM object.
func (d *DOM) Set(property string, value any) {
	d.jsValue.Set(property, unwrapValue(value))
}

// Call invokes a method of the DOM object with the given arguments and returns the result as new *DOM object.
func (d *DOM) Call(method string, args ...any) *DOM {
	return &DOM{
		jsValue: d.jsValue.Call(method, unwrapValues(args)...),
	}
}

// Await waits for a JavaScript promise to resolve or reject, returning the result as a new *DOM object or an error.
// It should be used in JavaScript functions that return a promise.
func (d *DOM) Await() (*DOM, error) {
	// Create channels to receive the result.
	valueChan := make(chan *DOM)
	errorChan := make(chan error)
	// Create JavaScript callback functions to handle the result.
	onFulfilled := onFulfilledCallback(valueChan)
	defer onFulfilled.Release()
	onRejected := onRejectedCallback(errorChan)
	defer onRejected.Release()
	// Call the JavaScript promise's "then" method with the callback functions.
	d.Call("then", onFulfilled, onRejected)
	// Wait for the result and return.
	select {
	case v := <-valueChan:
		return v, nil
	case e := <-errorChan:
		return nil, e
	}
}

// GetGlobal returns the representation of the JavaScript global object.
func GetGlobal() *DOM {
	return &DOM{
		jsValue: js.Global(),
	}
}

// Fetch performs an HTTP request and unmarshals the response into a Go struct of type T.
func Fetch[T any](url string) (*T, error) {
	// Use the browser's fetch API to perform the HTTP request.
	v, err := GetGlobal().Call("fetch", url).Await()
	if err != nil {
		return nil, err
	}
	// Extract the response body as a string.
	v, err = v.Call("text").Await()
	if err != nil {
		return nil, err
	}
	// Unmarshal the response into a Go struct of type T.
	var t T
	err = json.Unmarshal([]byte(v.String()), &t)
	if err != nil {
		return nil, err
	}
	// Return the result.
	return &t, nil
}

// unwrapValue unwraps a DOM value into a JavaScript value.
func unwrapValue(value any) any {
	v, ok := value.(*DOM)
	if ok {
		return v.jsValue
	}
	return value
}

// unwrapValues unwraps a slice of DOM values into a slice of JavaScript values.
func unwrapValues(values []any) []any {
	unwrappedValues := make([]any, len(values))
	for i, v := range values {
		unwrappedValues[i] = unwrapValue(v)
	}
	return unwrappedValues
}

// onFulfilledCallback returns a JavaScript function that is called when a promise is fulfilled.
func onFulfilledCallback(valueChan chan<- *DOM) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		valueChan <- &DOM{
			jsValue: args[0],
		}
		return nil
	})
}

// onRejectedCallback returns a JavaScript function that is called when a promise is rejected.
func onRejectedCallback(errorChan chan<- error) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		errorChan <- errors.New(args[0].String())
		return nil
	})
}
