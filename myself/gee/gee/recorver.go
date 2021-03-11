package gee

import "log"

func Recorver() HandleFunc {
	return func(c *Context) {
		defer func() {

			if err := recover(); err != nil {
				log.Println(err)
				c.W.Write([]byte("程序执行错误"))
			}
		}()
		// 重要这里要使程序在一个线程汇总
		// 可能defer中的recorver 需要在一个函数体中
		c.Next()
	}
}
