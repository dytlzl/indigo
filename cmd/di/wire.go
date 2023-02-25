//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/infra/repository"
	"github.com/dytlzl/indigo/pkg/usecase"
)

var wireSet = wire.NewSet(
	api.NewClient,
	repository.NewAPIInstanceRepository,
	repository.NewAPIFirewallRepository,
	repository.NewAPIOSRepository,
	repository.NewJSONPlanRepository,
	repository.NewAPISSHKeyRepository,
)

func InitializeInstanceUseCase(conf config.Config) usecase.InstanceUseCase {
	wire.Build(wireSet, usecase.NewInstanceUseCase)
	return nil
}

func InitializeOSUseCase(conf config.Config) usecase.OSUseCase {
	wire.Build(wireSet, usecase.NewOSUseCase)
	return nil
}

func InitializePlanUseCase(conf config.Config) usecase.PlanUseCase {
	wire.Build(wireSet, usecase.NewPlanUseCase)
	return nil
}

func InitializeSSHKeyUseCase(conf config.Config) usecase.SSHKeyUseCase {
	wire.Build(wireSet, usecase.NewSSHKeyUseCase)
	return nil
}

func InitializeFirewallUseCase(conf config.Config) usecase.FirewallUseCase {
	wire.Build(wireSet, usecase.NewFirewallUseCase)
	return nil
}
