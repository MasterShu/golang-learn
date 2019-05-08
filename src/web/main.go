package main

import (
	"net/http"
	"web/app"
)

func main() {
	server := http.Server{
		Addr: ":10015",
	}

	s := "/post/"
	http.HandleFunc(s, app.HandleRequest)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
