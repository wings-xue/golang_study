package main

import (
	"gee/engine"
	"net/http"
)

func main() {

	r := engine.New()
	r.Use(engine.Logger()) // global midlleware
	r.Static("/static", "readme.md")
	r.GET("/", func(c *engine.Context) {
		c.String(http.StatusOK, "<h1>Hello engine</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(engine.OnlyForV2()) // v2 group middleware

	v2.GET("/hello/:name", func(c *engine.Context) {
		// expect /hello/enginektutu

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param["name"], c.Path)
	})

	r.GET("/html", func(c *engine.Context) {
		c.HTML(http.StatusOK, "man")
	})
	r.Run(":9999")
}
