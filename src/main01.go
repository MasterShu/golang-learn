package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello Go WEb %s!", req.URL.Path[1:])
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	_ = http.ListenAndServe(":10010", nil)
}
