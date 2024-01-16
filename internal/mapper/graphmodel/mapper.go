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
		ID:        strconv.FormatInt(domainGroup.ID, 10),
		CreatorID: strconv.FormatInt(domainGroup.CreatorID, 10),
		Name:      domainGroup.GroupName,
	}
	return modelGroup
}

func MappingActivityDomainToGraphqlModel(domainActivity domain.Activity) *model.Activity {
	modelActivity := &model.Activity{
		ID:        strconv.FormatInt(domainActivity.ID, 10),
		GroupID:   strconv.FormatInt(domainActivity.GroupID, 10),
		CreatorID: strconv.FormatInt(domainActivity.CreatorID, 10),
		Name:      domainActivity.ActivityName,
	}
	return modelActivity
}
