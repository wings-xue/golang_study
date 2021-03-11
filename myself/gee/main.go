package main

import (
	"fmt"
	"gee/gee"
	"html"
)

func main() {

	// http.ListenAndServe(":8080", nil)

	router := gee.Default()

	router.GET("/bar/:name", func(c *gee.Context) {
		fmt.Fprintf(c.W, "Hello, %q", html.EscapeString(c.Param("name")))
	})

	router.GET("/panic", func(c *gee.Context) {
		a := make([]string, 0)
		fmt.Fprintf(c.W, "Hello, %q", a[100])
	})

	v1 := router.Group("/v1")
	v1.Use(gee.Auth())
	v1.GET("/bar/:name", func(c *gee.Context) {
		fmt.Fprintf(c.W, "Hello, %q", html.EscapeString(c.Param("name")))
	})
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
