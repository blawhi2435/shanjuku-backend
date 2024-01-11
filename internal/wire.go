//go:build wireinject
// +build wireinject
package internal

import (
	"github.com/blawhi2435/shanjuku-backend/graph/resolver"
	_authUsecase "github.com/blawhi2435/shanjuku-backend/internal/app/auth/usecase"
	_userUsecase "github.com/blawhi2435/shanjuku-backend/internal/app/user/usecase"
	_userPostgresRepo "github.com/blawhi2435/shanjuku-backend/internal/app/user/repository/postgres"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var resolverSet = wire.NewSet(resolver.ProvideResolver)
var usecaseSet = wire.NewSet(_userPostgresRepo.ProvideUserPostgresRepository, 
	_userUsecase.ProvideUserUsecase, _authUsecase.ProvideAuthUsecase)

func InitResolver(db *gorm.DB) *resolver.Resolver {
	wire.Build(resolverSet, usecaseSet)
	return nil
}
