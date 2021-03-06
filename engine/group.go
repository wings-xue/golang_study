package engine

type RouterGroup struct {
	prefix string
	middle []handleFunc
	router *Router
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	routerGroup := &RouterGroup{
		prefix: prefix,
		router: g.router,
	}
	g.router.routerGroups = append(g.router.routerGroups, routerGroup)
	return routerGroup
}

func (g *RouterGroup) GET(addr string, handleFunc handleFunc) {
	g.router.addRouter("GET", g.prefix+addr, handleFunc)
}

func (g *RouterGroup) POST(addr string, handleFunc handleFunc) {
	g.router.addRouter("POST", g.prefix+addr, handleFunc)
}

func (g *RouterGroup) Use(middle ...handleFunc) {
	g.middle = append(g.middle, middle...)
}
