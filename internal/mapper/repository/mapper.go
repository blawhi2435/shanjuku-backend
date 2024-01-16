package repository

import (
	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
)

func MappingUserDomainToSchema(domainUser *domain.User) (schemaUser *postgres.User) {
	schemaUser = &postgres.User{
		ID:       domainUser.ID,
		Account:  domainUser.Account,
		Password: domainUser.Password,
		Name:     domainUser.Name,
		Avatar:   domainUser.Avatar,
	}
	return
}

func MappingUserSchemaToDomain(schemaUser *postgres.User) (domainUser domain.User) {
	domainUser = domain.User{
		ID:       schemaUser.ID,
		Account:  schemaUser.Account,
		Password: schemaUser.Password,
		Name:     schemaUser.Name,
		Avatar:   schemaUser.Avatar,
	}
	return
}

func MappingGroupDomainToSchema(domainGroup *domain.Group) (schemaGroup *postgres.Group) {
	schemaGroup = &postgres.Group{
		ID:        domainGroup.ID,
		CreatorID: domainGroup.CreatorID,
		GroupName: domainGroup.GroupName,
	}
	return
}

func MappingGroupSchemaGroupToDomain(schemaGroup *postgres.Group) (domainGroup domain.Group) {
	domainGroup = domain.Group{
		ID:        schemaGroup.ID,
		CreatorID: schemaGroup.CreatorID,
		GroupName: schemaGroup.GroupName,
	}
	return
}

func MappingActivityDomainToSchema(domainActivity *domain.Activity) (schemaActivity *postgres.Activity) {
	schemaActivity = &postgres.Activity{
		ID:           domainActivity.ID,
		CreatorID:    domainActivity.CreatorID,
		GroupID:      domainActivity.GroupID,
		ActivityName: domainActivity.ActivityName,
	}
	return
}

func MappingActivitySchemaToDomain(schemaActivity *postgres.Activity) (domainActivity domain.Activity) {
	domainActivity = domain.Activity{
		ID:           schemaActivity.ID,
		CreatorID:    schemaActivity.CreatorID,
		GroupID:      schemaActivity.GroupID,
		ActivityName: schemaActivity.ActivityName,
	}
	return
}
