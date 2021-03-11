package gee

type HandleFunc func(*Context)

type Router struct {
	handle map[string]HandleFunc
	root   map[string]*Node
}

func (r *Router) addRouter(method, pattern string, handle HandleFunc) {
	r.handle[method+"-"+pattern] = handle
	if n, ok := r.root[method]; ok {
		n.insert(pattern, parsePath(pattern), 0)
	} else {
		n := NewRoot()
		n.insert(pattern, parsePath(pattern), 0)
		r.root[method] = n

	}
}

func (r *Router) getRouter(method, path string) (string, HandleFunc) {
	if n, ok := r.root[method]; ok {
		node := n.search(path, parsePath(path), 0)
		return node.pattern, r.handle[method+"-"+node.pattern]
	}
	return "", nil
}
