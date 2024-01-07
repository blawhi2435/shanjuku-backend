package service

import "gorm.io/gorm"

type PostgreService struct {
	DB *gorm.DB
}