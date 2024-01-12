package mapper

import (
	"strconv"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/graph/model"
)

func MappingUserDomainToModel(domainUser domain.User) (*model.User) {
	modelUser := &model.User{
		ID:       strconv.FormatInt(domainUser.ID, 10),
		Account:  domainUser.Account,
		Name:     domainUser.Name,
		Avatar:   domainUser.Avatar,
	}
	return modelUser
}