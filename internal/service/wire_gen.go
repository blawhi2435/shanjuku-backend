// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"fmt"
	"github.com/blawhi2435/shanjuku-backend/enviroment"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitService() (*Service, error) {
	ginService := ProvideGinService()
	postgresService, err := ProvidePostgreService()
	if err != nil {
		return nil, err
	}
	service := ProvideService(ginService, postgresService)
	return service, nil
}

// wire.go:

var serviceSet = wire.NewSet(ProvideGinService, ProvidePostgreService, ProvideService)

func ProvideGinService() *GinService {
	r := gin.Default()

	return &GinService{Engine: r}
}

func ProvidePostgreService() (*PostgresService, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", enviroment.Setting.
		Postgres.Host, enviroment.Setting.Postgres.User, enviroment.Setting.Postgres.Password, enviroment.Setting.
		Postgres.Database, enviroment.Setting.Postgres.Port, enviroment.Setting.Postgres.TimeZone)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &PostgresService{DB: db}, nil
}

func ProvideService(g *GinService, p *PostgresService) *Service {
	return &Service{GinService: g, PostgresService: p}
}
