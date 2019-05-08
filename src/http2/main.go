package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "hello go!")
}
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:10012",
		Handler: &handler,
	}
	_ = http2.ConfigureServer(&server, &http2.Server{})
	_ = server.ListenAndServeTLS("cert.pem", "key.pem")

}
