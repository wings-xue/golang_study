package main

import (
	"gee/engine"
	"io"
	"net/http"
)

func searchHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "search succeed")
}

func saveHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "save succeed")
}

func main() {

	r := engine.New()
	r.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *engine.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *engine.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
