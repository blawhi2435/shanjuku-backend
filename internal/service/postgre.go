package service

import "gorm.io/gorm"

type PostgresService struct {
	DB *gorm.DB
}
