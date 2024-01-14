package postgres

import (
	"errors"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

func ProvideDBRepository(db *gorm.DB) domain.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) error) (err error) {
	tx := r.db.Begin()
	if !errors.Is(tx.Error, nil) {
		return tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if !errors.Is(err, nil) {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)
	return err
}