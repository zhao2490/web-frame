package middleware

import (
	"log"

	"github.com/zhao2490/web-frame/gee"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("[%s] Requst [%s]%s", c.Req.RemoteAddr, c.Req.Method, c.Req.URL.Path)
		c.Next()
	}
}
