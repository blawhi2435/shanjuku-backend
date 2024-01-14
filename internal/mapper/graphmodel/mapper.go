package graphmodel

import (
	"strconv"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/graph/model"
)

func MappingUserDomainToGraphqlModel(domainUser domain.User) *model.User {
	modelUser := &model.User{
		ID:      strconv.FormatInt(domainUser.ID, 10),
		Account: domainUser.Account,
		Name:    domainUser.Name,
		Avatar:  domainUser.Avatar,
	}
	return modelUser
}

func MappingGroupDomainToGraphqlModel(domainGroup domain.Group) *model.Group {
	modelGroup := &model.Group{
		ID:   strconv.FormatInt(domainGroup.ID, 10),
		Name: domainGroup.Name,
	}
	return modelGroup
}