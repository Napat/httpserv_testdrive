package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/":
		io.WriteString(res, "<strong>I am Root<strong>")
	case "/ant":
		io.WriteString(res, "<strong>And An Ant<strong>")
	case "/bird":
		io.WriteString(res, "<strong>Big Bird<strong>")
	}
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
