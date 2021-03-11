package gee

import "log"

func Logger() HandleFunc {
	return func(c *Context) {
		log.Printf("请求%s \n", c.Req.URL)

	}
}
