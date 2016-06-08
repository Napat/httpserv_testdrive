package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, GetOutboundIP())

	switch req.URL.Path {
	case "/":
		io.WriteString(res, "<p><strong>I AM gROOT<strong></p>")
	case "/ant":
		io.WriteString(res, "<p><strong>And Ann Ant<strong></p>")
	case "/bird":
		io.WriteString(res, "<p><strong>Black Bird<strong></p>")
	}
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
