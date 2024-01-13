//go:build wireinject
// +build wireinject
package internal

import (
	"github.com/blawhi2435/shanjuku-backend/graph/resolver"
	_authUsecase "github.com/blawhi2435/shanjuku-backend/internal/app/auth/usecase"
	_userPostgresRepo "github.com/blawhi2435/shanjuku-backend/internal/app/user/repository/postgres"
	_groupPostgresRepo "github.com/blawhi2435/shanjuku-backend/internal/app/group/repository/postgres"
	_userUsecase "github.com/blawhi2435/shanjuku-backend/internal/app/user/usecase"
	_groupUsecase "github.com/blawhi2435/shanjuku-backend/internal/app/group/usecase"
	"github.com/blawhi2435/shanjuku-backend/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var resolverSet = wire.NewSet(resolver.ProvideResolver)
var usecaseSet = wire.NewSet(_userPostgresRepo.ProvideUserPostgresRepository, 
	_groupPostgresRepo.ProvideGroupPostgresRepository, _userUsecase.ProvideUserUsecase, 
	_authUsecase.ProvideAuthUsecase, _groupUsecase.ProvideGroupUsecase)

func InitResolver(db *gorm.DB, logger *service.LoggerService) *resolver.Resolver {
	wire.Build(resolverSet, usecaseSet)
	return nil
}
