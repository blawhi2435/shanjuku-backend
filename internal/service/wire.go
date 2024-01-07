//go:build wireinject
// +build wireinject

package service

import (
	"fmt"

	"github.com/blawhi2435/shanjuku-backend/enviroment"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var serviceSet = wire.NewSet(ProvideGinService, ProvidePostgreService, ProvideService)

func ProvideGinService() *GinService {
	r := gin.Default()

	return &GinService{Engine: r}
}

func ProvidePostgreService() (*PostgreService, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		enviroment.Setting.Postgre.Host, enviroment.Setting.Postgre.User, enviroment.Setting.Postgre.Password, 
		enviroment.Setting.Postgre.Database, enviroment.Setting.Postgre.Port, enviroment.Setting.Postgre.TimeZone)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &PostgreService{DB: db}, nil
}

func ProvideService(g *GinService, p *PostgreService) *Service {
	return &Service{GinService: g, PostgreService: p}
}

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
