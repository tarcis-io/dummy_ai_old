package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/pkg/env"
)

var (
	serverTemplate = template.Must(template.ParseFiles("./web/server/server.html"))
)

func Start() {

	serveFile("/wasm_exec.js" /*   */, "./web/static/lib/wasm/wasm_exec.js")
	serveFile("/wasm_start.js" /*  */, "./web/static/lib/wasm/wasm_start.js")
	serveFile("/error_404.wasm" /* */, "./web/static/wasm/error_404.wasm")
	serveFile("/index.wasm" /*     */, "./web/static/wasm/index.wasm")

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func servePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			error404(responseWriter)
			return
		}

		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}

func error404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}
