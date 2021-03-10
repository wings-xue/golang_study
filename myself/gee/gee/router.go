package gee

import "net/http"

type HandleFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	handle map[string]HandleFunc
	root   *Node
}

func Default() *Router {
	return &Router{
		handle: make(map[string]HandleFunc),
		root:   NewRoot(),
	}
}

func (r *Router) addRouter(method, pattern string, handle HandleFunc) {
	r.handle[method+"-"+pattern] = handle
	r.root.insert(pattern, parsePath(pattern), 0)
}

func (r *Router) GET(pattern string, handle HandleFunc) {
	r.addRouter("GET", pattern, handle)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	node := r.root.search(path, parsePath(path), 0)
	if node == nil {
		w.Write([]byte("400 not find"))
		return
	}

	handle, ok := r.handle[method+"-"+node.pattern]
	if !ok {
		w.Write([]byte("400 not find"))
		return
	}

	handle(w, req)

}

func (r *Router) Run() {

	http.ListenAndServe(":8080", r)

}
