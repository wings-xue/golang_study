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
	roots  map[string]*GoodNode
}

func New() *Router {
	handle := make(map[string]handleFunc)
	roots := make(map[string]*GoodNode)
	return &Router{
		handle: handle,
		roots:  roots,
	}
}

func (e *Router) insert(method, addr string) {
	if root, ok := e.roots[method]; ok {
		root.insert(addr, parsePath(addr), 0)
	} else {
		newNode := NewRoot()
		newNode.insert(addr, parsePath(addr), 0)
		e.roots[method] = newNode
	}
}

func (e *Router) search(method, addr string) (handleFunc, []string, []string) {
	if root, ok := e.roots[method]; ok {
		path := parsePath(addr)
		node := root.search(addr, path, 0)
		if node == nil {
			return nil, []string{}, []string{}
		}
		pattern := node.pattern
		return e.handle[routerKey(method, pattern)], path, parsePath(pattern)

	}
	return nil, []string{}, []string{}
}

func (e *Router) addRouter(method, addr string, handleFunc handleFunc) {

	e.insert(method, addr)
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
	c := Context{
		Writer: w,
		Req:    req,
		Param:  map[string]string{},
	}
	handleFunc, paths, patterns := e.search(method, addr)
	if handleFunc == nil {
		io.WriteString(w, "404 NOT FOUND\n")
		return
	}
	c.FindParam(paths, patterns)

	handleFunc(&c)
}

func (e *Router) Run(addr string) {
	http.ListenAndServe(addr, e)
}
