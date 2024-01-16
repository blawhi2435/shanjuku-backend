package resolver

import (
	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger          *service.LoggerService
	AuthUsecasse    domain.AuthUsecase
	UserUsecase     domain.UserUsecase
	GroupUsecase    domain.GroupUsecase
	ActivityUsecase domain.ActivityUsecase
}

func ProvideResolver(
	logger *service.LoggerService,
	authUsecase domain.AuthUsecase,
	userUsecase domain.UserUsecase,
	groupUsecase domain.GroupUsecase,
	activityUsecase domain.ActivityUsecase,
) *Resolver {
	return &Resolver{
		logger:          logger,
		AuthUsecasse:    authUsecase,
		UserUsecase:     userUsecase,
		GroupUsecase:    groupUsecase,
		ActivityUsecase: activityUsecase,
	}
}
