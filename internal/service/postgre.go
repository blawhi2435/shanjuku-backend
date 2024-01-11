package service

import (
	"fmt"

	"github.com/blawhi2435/shanjuku-backend/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresService struct {
	DB *gorm.DB
}

func ProvidePostgreService() (*PostgresService, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		environment.Setting.Postgres.Host, environment.Setting.Postgres.User, environment.Setting.Postgres.Password,
		environment.Setting.Postgres.Database, environment.Setting.Postgres.Port, environment.Setting.Postgres.TimeZone)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &PostgresService{DB: db}, nil
}
