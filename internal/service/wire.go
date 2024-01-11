//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(ProvideGinService, ProvidePostgreService, ProvideLogger, ProvideService)

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
