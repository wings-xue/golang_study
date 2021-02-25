package engine

import (
	"io"
	"net/http"
)

// 处理路由

type handleFunc func(c *Context)

func routerKey(method, addr string) string {
	return method + "-" + addr
}

type Router struct {
	handle map[string]handleFunc
}

func New() *Router {
	handle := make(map[string]handleFunc)
	return &Router{handle: handle}
}

func (e *Router) addRouter(method, addr string, handleFunc handleFunc) {
	e.handle[routerKey(method, addr)] = handleFunc
}

func (e *Router) GET(addr string, handleFunc handleFunc) {
	e.addRouter("GET", addr, handleFunc)
}

func (e *Router) POST(addr string, handleFunc handleFunc) {
	e.addRouter("POST", addr, handleFunc)
}

func (e *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	addr := req.URL.Path
	method := req.Method
	if handleFunc, ok := e.handle[routerKey(method, addr)]; ok {
		c := Context{
			Writer: w,
			Req:    req,
		}
		handleFunc(&c)
		return
	}
	io.WriteString(w, "404 NOT FOUND\n")

}

func (e *Router) Run(addr string) {
	http.ListenAndServe(addr, e)
}
