package resolver

import "github.com/blawhi2435/shanjuku-backend/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	AuthUsecasse domain.AuthUsecase
}

func ProvideResolver(authUsecase domain.AuthUsecase) *Resolver {
	return &Resolver{AuthUsecasse: authUsecase}
}