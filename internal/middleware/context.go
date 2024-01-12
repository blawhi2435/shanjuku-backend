package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

var GinContextKey = "ginContextKey"

func  GinContextToContext() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}