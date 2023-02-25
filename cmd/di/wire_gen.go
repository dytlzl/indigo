// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/infra/repository"
	"github.com/dytlzl/indigo/pkg/usecase"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeInstanceUseCase(conf config.Config) usecase.InstanceUseCase {
	client := api.NewClient(conf)
	instanceRepository := repository.NewAPIInstanceRepository(client)
	instanceUseCase := usecase.NewInstanceUseCase(instanceRepository)
	return instanceUseCase
}

func InitializeOSUseCase(conf config.Config) usecase.OSUseCase {
	client := api.NewClient(conf)
	osRepository := repository.NewAPIOSRepository(client)
	osUseCase := usecase.NewOSUseCase(osRepository)
	return osUseCase
}

func InitializePlanUseCase(conf config.Config) usecase.PlanUseCase {
	planRepository := repository.NewJSONPlanRepository()
	planUseCase := usecase.NewPlanUseCase(planRepository)
	return planUseCase
}

func InitializeSSHKeyUseCase(conf config.Config) usecase.SSHKeyUseCase {
	client := api.NewClient(conf)
	sshKeyRepository := repository.NewAPISSHKeyRepository(client)
	sshKeyUseCase := usecase.NewSSHKeyUseCase(sshKeyRepository)
	return sshKeyUseCase
}

func InitializeFirewallUseCase(conf config.Config) usecase.FirewallUseCase {
	client := api.NewClient(conf)
	firewallRepository := repository.NewAPIFirewallRepository(client)
	instanceRepository := repository.NewAPIInstanceRepository(client)
	firewallUseCase := usecase.NewFirewallUseCase(firewallRepository, instanceRepository)
	return firewallUseCase
}

// wire.go:

var wireSet = wire.NewSet(api.NewClient, repository.NewAPIInstanceRepository, repository.NewAPIFirewallRepository, repository.NewAPIOSRepository, repository.NewJSONPlanRepository, repository.NewAPISSHKeyRepository)
