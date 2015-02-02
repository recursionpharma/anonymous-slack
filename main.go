package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"
)

var port int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func init() {
	flag.IntVar(&port, "port", 5000, "HTTP server port")
	flag.Parse()
}
