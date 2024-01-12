package service

import (
	"github.com/blawhi2435/shanjuku-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type GinService struct {
	Engine *gin.Engine
}

func ProvideGinService() *GinService {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.GinContextToContext())

	return &GinService{Engine: r}
}