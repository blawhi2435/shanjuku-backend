package middleware

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/internal/ctxtool"
	"github.com/gin-gonic/gin"
)

func GinContextToContext() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ctxtool.GinCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
