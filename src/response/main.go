package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, _ *http.Request) {
	str := `
<html>
<head><title>My respone test</title></head>
<body>
<h2> Hello Go</h2>
</body>
</html>
`
	_, _ = w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = fmt.Fprintf(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Location", "https://bing.com")
	w.WriteHeader(http.StatusFound)
}

func jsonExample(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type post struct {
		User   string
		Thread []string
	}
	postData := &post{
		User:   "MasterShu",
		Thread: []string{"hello", "world"},
	}
	jsonString, _ := json.Marshal(postData)
	_, _ = w.Write(jsonString)
}

func setCookies(w http.ResponseWriter, _ *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "My cookie test",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second",
		Value:    "Another cookie test",
		HttpOnly: true,
	}
	w.Header().Set("set-Cookie", c1.String())
	w.Header().Add("set-Cookie", c2.String())
	// Another styles
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookies(w http.ResponseWriter, req *http.Request) {
	h := req.Header["Cookie"]
	_, _ = fmt.Fprintln(w, h)
}

func getCookies2(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("first")
	if err != nil {
		_, _ = fmt.Fprintln(w, "Can not get first")
	}
	_, _ = fmt.Fprintln(w, c1)
	cs := req.Cookies()
	_, _ = fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: ":10013",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/cookie", setCookies)
	http.HandleFunc("/cookie_get", getCookies)
	http.HandleFunc("/cookie_get2", getCookies2)
	_ = server.ListenAndServe()

}
