package engine

import (
	"net/http"
	"path"
)

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

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) handleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		// file := c.Param["filepath"]
		// // Check if file exists and/or if we have permission to access it
		// if _, err := fs.Open("D:/code/wing-xue/golang_study/" + file); err != nil {
		// 	c.Status(http.StatusNotFound)
		// 	return
		// }

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}
