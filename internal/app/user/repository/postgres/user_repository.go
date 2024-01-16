package postgres

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/mapper/repository"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func ProvideUserPostgresRepository(db *gorm.DB) domain.UserPostgreRepository {
	return &userPostgresRepository{db}
}

func (u *userPostgresRepository) InsertNewUser(ctx context.Context, user *domain.User) error {

	schemaUser := repository.MappingUserDomainToSchema(user)
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

	if res.Error == gorm.ErrRecordNotFound {
		return domain.User{}, cerror.ErrUserNotExist
	}

	return repository.MappingUserSchemaToDomain(&schemaUser), res.Error
}

func (u *userPostgresRepository) QueryByID(ctx context.Context, id int64) (domain.User, error) {
	var schemaUser postgres.User
	orm := u.db.Model(&schemaUser)
	res := orm.Select("id", "account", "password", "name", "avatar").
		Where("id = ?", id).
		Take(&schemaUser)

	if res.Error == gorm.ErrRecordNotFound {
		return domain.User{}, cerror.ErrUserNotExist
	}

	return repository.MappingUserSchemaToDomain(&schemaUser), res.Error
}

func (u *userPostgresRepository) QueryByIDs(ctx context.Context, ids []int64) ([]domain.User, error) {
	var schemaUsers []postgres.User
	var domainUsers []domain.User

	orm := u.db.Model(&schemaUsers)
	res := orm.Select("id", "account", "password", "name", "avatar").
		Where("id IN ?", ids).
		Find(&schemaUsers)

	for _, schemaUser := range schemaUsers {
		domainUsers = append(domainUsers, repository.MappingUserSchemaToDomain(&schemaUser))
	}

	return domainUsers, res.Error
}

func (u *userPostgresRepository) QueryAssociationGroups(ctx context.Context, 
	userID int64) ([]domain.Group, error) {

	var schemaGroups []*postgres.Group
	schemaUser := postgres.User{
		ID: userID,
	}

	orm := u.db.Model(&schemaUser)
	orm.Association("Groups").
		Find(&schemaGroups)

	var groups []domain.Group
	for _, schemaGroup := range schemaGroups {
		groups = append(groups, repository.MappingGroupSchemaGroupToDomain(schemaGroup))
	}

	return groups, orm.Error
}
