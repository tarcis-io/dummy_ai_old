package dom

import (
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
	return nil, nil
}

func GetGlobal() *DOM {
	return &DOM{
		jsValue: js.Global(),
	}
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
