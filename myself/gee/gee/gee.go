package gee

import (
	"net/http"
	"strings"
)

type GroupRouter struct {
	Prefix string
	*Engine
	Middle []HandleFunc
}
type Engine struct {
	*Router
	groups []*GroupRouter
	*GroupRouter
}

func Default() *Engine {
	router := &Router{
		handle: make(map[string]HandleFunc),
		root:   make(map[string]*Node),
	}
	e := &Engine{
		Router: router,
		groups: make([]*GroupRouter, 0),
	}
	e.GroupRouter = &GroupRouter{
		Engine: e,
	}
	e.Use(Recorver())
	e.Use(Logger())

	return e
}

func (g *GroupRouter) Group(prefix string) *GroupRouter {

	group := &GroupRouter{
		Prefix: prefix,
		Engine: g.Engine,
		Middle: make([]HandleFunc, 0),
	}
	g.Engine.groups = append(g.Engine.groups, group)
	return group
}

func (g *GroupRouter) GET(pattern string, handle HandleFunc) {
	g.Router.addRouter("GET", g.Prefix+pattern, handle)
}

func (g *GroupRouter) Use(f HandleFunc) {
	g.Middle = append(g.Middle, f)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	Middle := make([]HandleFunc, 0)

	Middle = append(Middle, e.GroupRouter.Middle...)

	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.Prefix) {
			Middle = append(Middle, group.Middle...)
		}
	}
	c := NewContext(req, w)
	c.Middle = Middle

	pattern, handle := e.Router.getRouter(req.Method, req.URL.Path)
	c.Middle = append(c.Middle, handle)
	c.param(pattern, req.URL.Path)
	c.Next()
}

func (e *Engine) Run() {

	http.ListenAndServe(":8080", e)

}
