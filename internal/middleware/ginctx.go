package middleware

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/contextkey"
	"github.com/gin-gonic/gin"
)

func GinContextToContext() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextkey.GinCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(contextkey.GinCtxKey)
	if ginContext == nil {
		return nil, cerror.ErrGetContextFailed
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, cerror.ErrGetContextFailed
	}

	return gc, nil
}
