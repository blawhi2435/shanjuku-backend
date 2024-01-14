package ctxtool

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/gin-gonic/gin"
)

func SetTokenToContext(ctx context.Context) (context.Context, error) {
	gctx, err := ginContextFromContext(ctx)
	if err != nil {
		return ctx, cerror.GetGQLError(ctx, err)
	}

	token := gctx.GetString(TokenCtxKey)
	ctx = context.WithValue(ctx, TokenCtxKey, token)

	return ctx, nil
}

func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinCtxKey)
	if ginContext == nil {
		return nil, cerror.ErrGetContextFailed
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, cerror.ErrGetContextFailed
	}

	return gc, nil
}