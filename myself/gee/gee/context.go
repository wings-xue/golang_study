package gee

import (
	"net/http"
	"strings"
)

type Context struct {
	Req    *http.Request
	W      http.ResponseWriter
	Params map[string]string
	Middle []HandleFunc
	index  int
}

func NewContext(req *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		W:      w,
		Params: make(map[string]string),
		Middle: make([]HandleFunc, 0),
		index:  -1,
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) param(pattern, path string) {
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")
	for i, v := range patterns {
		switch {
		case strings.HasPrefix(v, ":"):
			c.Params[v[1:]] = paths[i]
		case strings.HasPrefix(v, "*"):
			c.Params[v[1:]] = strings.Join(paths[i:], "/")
			break

		}

	}

}

func (c *Context) Next() {
	c.index = c.index + 1
	for ; c.index < len(c.Middle); c.index++ {
		c.Middle[c.index](c)
	}
}
