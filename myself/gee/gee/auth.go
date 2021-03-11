package gee

import "log"

func Auth() HandleFunc {
	return func(c *Context) {
		log.Printf("v1 正在登陆...")
	}
}
