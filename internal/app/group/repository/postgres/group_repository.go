package postgres

import (
	"context"
	"time"

	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/mapper/repository"
	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type groupName struct {
	Name string `structs:"name"`
}

type groupPostgresRepository struct {
	db *gorm.DB
}

func ProvideGroupPostgresRepository(db *gorm.DB) domain.GroupPostgreRepository {
	return &groupPostgresRepository{db}
}

func (g *groupPostgresRepository) InsertGroup(ctx context.Context, group *domain.Group) error {
	user := postgres.User{
		ID: group.CreatorID,
	}

	schemaGroup := repository.MappingGroupDomainToSchema(group)
	schemaGroup.CreatedDate = time.Now()

	orm := g.db.Model(&user)
	orm.Association("Groups").Append(schemaGroup)

	return orm.Error
}

func (g *groupPostgresRepository) QueryByID(ctx context.Context, groupID int64) (domain.Group, error) {
	var schemaGroup postgres.Group

	orm := g.db.Model(&schemaGroup)
	orm.Where("id = ?", groupID).
		Take(&schemaGroup)

	return repository.MappingGroupSchemaGroupToDomain(&schemaGroup), orm.Error
}

func (g *groupPostgresRepository) QueryAssociationUsers(ctx context.Context,
	groupID int64) ([]domain.User, error) {

	var schemaUsers []*postgres.User
	schemaGroup := postgres.Group{
		ID: groupID,
	}

	orm := g.db.Model(&schemaGroup)
	orm.Association("Users").
		Find(&schemaUsers)

	var users []domain.User
	for _, user := range schemaUsers {
		users = append(users, repository.MappingUserSchemaToDomain(user))
	}

	return users, orm.Error
}

func (g *groupPostgresRepository) UpdateGroup(ctx context.Context, group *domain.Group) (int64, error) {

	updateColumn := groupName{
		Name: group.Name,
	}

	res := g.db.Model(&postgres.Group{}).
		Where("group_id = ?", group.ID).
		Updates(structs.Map(updateColumn))

	return res.RowsAffected, res.Error
}

func (g *groupPostgresRepository) DeleteGroup(ctx context.Context, groupID int64) error {

	res := g.db.Delete(&postgres.Group{}, groupID)

	return res.Error
}

func (g *groupPostgresRepository) AddUserToGroup(ctx context.Context,
	groupID int64, userID []int64) error {

	group := postgres.Group{
		ID: groupID,
	}

	var users []*postgres.User
	for _, id := range userID {
		users = append(users, &postgres.User{
			ID: id,
		})
	}

	orm := g.db.Model(&group)
	err := orm.Association("Users").Append(&users)

	return err
}

func (g *groupPostgresRepository) RemoveUserFromGroup(ctx context.Context, groupID, userID int64) error {
	group := postgres.Group{
		ID: groupID,
	}

	user := postgres.User{
		ID: userID,
	}

	orm := g.db.Model(&group)
	err := orm.Association("Users").Delete(&user)

	return err
}

func (g *groupPostgresRepository) ClearUserInGroup(ctx context.Context, groupID int64) error {
	group := postgres.Group{
		ID: groupID,
	}

	orm := g.db.Model(&group)
	err := orm.Association("Users").Clear()

	return err
}