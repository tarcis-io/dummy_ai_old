package server

import (
	"text/template"
)

var (
	httpTemplate = template.Must(template.ParseFiles("template.html"))
)
