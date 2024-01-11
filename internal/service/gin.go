package service

import "github.com/gin-gonic/gin"

type GinService struct {
	Engine *gin.Engine
}

func ProvideGinService() *GinService {
	r := gin.Default()

	return &GinService{Engine: r}
}