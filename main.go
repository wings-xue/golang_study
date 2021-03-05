package main

import (
	"gee/engine"
	"net/http"
)

func main() {

	r := engine.New()
	r.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello/:name", func(c *engine.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param["name"], c.Param["name"])
	})

	r.POST("/login", func(c *engine.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8881")
}
