package postgres

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func ProvideUserPostgresRepository(db *gorm.DB) domain.UserPostgreRepository {
	return &userPostgresRepository{db}
}

func (u *userPostgresRepository) InsertNewUser(ctx context.Context, user *domain.User) error {

	schemaUser := mappingDomainToSchema(user)
	orm := u.db.Model(&schemaUser)
	res := orm.Select("id", "account", "password").
		Create(&schemaUser)

	return res.Error
}

func (u *userPostgresRepository) QueryByAccount(ctx context.Context, account string) (domain.User, error) {
	var schemaUser postgres.User
	orm := u.db.Model(&schemaUser)
	res := orm.Select("id", "account", "password", "name", "avatar").
		Where("account = ?", account).
		Take(&schemaUser)

	return mappingSchemaToDomain(&schemaUser), res.Error
}
