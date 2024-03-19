package middleware

import (
	"context"
	"github.com/blocktock/go-pkg/tracex"
	"github.com/gin-gonic/gin"
)

// Get or set trace_id in request context
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = tracex.NewTraceID()
		}

		ctx := context.WithValue(c.Request.Context(), tracex.TraceIDCtx{}, traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
