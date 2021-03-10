package main

import (
	"fmt"
	"gee/gee"
	"html"
	"net/http"
)

func main() {

	// http.ListenAndServe(":8080", nil)

	router := gee.Default()

	router.GET("/bar/:name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
