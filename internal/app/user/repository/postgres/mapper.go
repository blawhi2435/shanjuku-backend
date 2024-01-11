package postgres

import (
	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
)

func mappingDomainToSchema(domainUser *domain.User) (schemaUser *postgres.User) {
	schemaUser = &postgres.User{
		ID:       domainUser.ID,
		Account:  domainUser.Account,
		Password: domainUser.Password,
		Name:     domainUser.Name,
		Avatar:   domainUser.Avatar,
	}
	return
}

func mappingSchemaToDomain(schemaUser *postgres.User) (domainUser domain.User) {
	domainUser = domain.User{
		ID:       schemaUser.ID,
		Account:  schemaUser.Account,
		Password: schemaUser.Password,
		Name:     schemaUser.Name,
		Avatar:   schemaUser.Avatar,
	}
	return
}