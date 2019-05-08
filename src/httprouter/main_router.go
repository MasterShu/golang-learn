package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	_, _ = fmt.Fprintf(w, "hello world %s!", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    ":10011",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
