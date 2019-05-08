package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, _ *http.Request) {
	msg := []byte("Hello go")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			_, _ = fmt.Fprintf(w, "Not message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		_, _ = fmt.Fprintln(w, string(val))
	}

}

func main() {

	server := http.Server{
		Addr: ":10014",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	_ = server.ListenAndServe()
}
