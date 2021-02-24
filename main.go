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
	e := engine.New()
	e.GET("search", searchHandle)
	e.POST("save", saveHandle)

	e.Run(":8888")
}
