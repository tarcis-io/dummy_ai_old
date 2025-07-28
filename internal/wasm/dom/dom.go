package dom

import (
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

func GetGlobal() *DOM {
	return &DOM{
		jsObject: js.Global(),
	}
}
