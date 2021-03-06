package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 处理请求和返回结果
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
	Param      map[string]string
	Index      int
	Middle     []handleFunc
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, val string) {
	c.Writer.Header().Add(key, val)
}
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) FindParam(paths, patterns []string) {
	for i, part := range patterns {
		if part[0] == ':' {
			c.Param[part[1:]] = paths[i]
		}
		if part[0] == '*' {
			c.Param[part[1:]] = strings.Join(paths[i:], "/")
			break
		}

	}
}

// 实现next，
func (c *Context) Next() {
	c.Index = c.Index + 1
	for ; c.Index < len(c.Middle); c.Index++ {
		c.Middle[c.Index](c)
	}
}
