package engine

import (
	"io"
	"net/http"
)

type handleFunc func(w http.ResponseWriter, req *http.Request)

func routerKey(method, addr string) string {
	return method + "-" + addr
}

type engine struct {
	handle map[string]handleFunc
}

func New() *engine {
	handle := make(map[string]handleFunc)
	return &engine{handle: handle}
}

func (e *engine) addRouter(method, addr string, handleFunc handleFunc) {
	e.handle[routerKey(method, addr)] = handleFunc
}

func (e *engine) GET(addr string, handleFunc handleFunc) {
	e.addRouter("GET", addr, handleFunc)
}

func (e *engine) POST(addr string, handleFunc handleFunc) {
	e.addRouter("POST", addr, handleFunc)
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	addr := req.URL.Path
	method := req.Method
	if handleFunc, ok := e.handle[routerKey(method, addr)]; ok {
		handleFunc(w, req)
		return
	}
	io.WriteString(w, "404 NOT FOUND\n")

}

func (e *engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
