package middleware

import (
	"github.com/blocktock/go-pkg/ginx"
	"github.com/gin-gonic/gin"
)

func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NoMethodMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginx.ResJSON(c, 404, "Not found")
		c.Abort()
	}
}

func NoRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginx.ResJSON(c, 404, "Not found")
		c.Abort()
	}
}

type SkipperFunc func(*gin.Context) bool

func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}
