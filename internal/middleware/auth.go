package middleware

import (
	"net/http"
	"strings"

	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/gin-gonic/gin"
)

func  AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		// Allow unauthenticated users in 
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusOK, cerror.GetGQLError(c, cerror.ErrTokenExpired))
			return
		}

		token := parts[1]
		// parts[0] is Bearer, parts[1] is token.
		c.Set("token", token)
		c.Next()
	}
}