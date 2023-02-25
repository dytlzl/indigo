package usecase

import (
	"context"
	"fmt"
	"sort"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printutil"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type InstanceRepository interface {
	List(ctx context.Context) ([]domain.Instance, error)
	Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error
	UpdateStatus(ctx context.Context, id int, status string) error
}

type InstanceUseCase interface {
	List(ctx context.Context) error
	Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error
	Start(ctx context.Context, name string) error
	Stop(ctx context.Context, name string) error
	ForceStop(ctx context.Context, name string) error
	Delete(ctx context.Context, name string) error
}

type instanceUseCase struct {
	instanceRepository InstanceRepository
}

func NewInstanceUseCase(i InstanceRepository) InstanceUseCase {
	return &instanceUseCase{
		instanceRepository: i,
	}
}

func (u *instanceUseCase) List(ctx context.Context) error {
	instances, err := u.instanceRepository.List(ctx)
	if err != nil {
		return err
	}
	sort.Slice(instances, func(i, j int) bool { return instances[i].Name < instances[j].Name })
	printutil.PrintTable(instances)
	return nil
}

func (u *instanceUseCase) Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error {
	return u.instanceRepository.Create(ctx, name, planID, osID, regionID, sshKeyID)
}

func (u *instanceUseCase) Start(ctx context.Context, name string) error {
	id, err := u.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = u.instanceRepository.UpdateStatus(ctx, id, "start")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" started\n", name)
	return nil
}

func (u *instanceUseCase) Stop(ctx context.Context, name string) error {
	id, err := u.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = u.instanceRepository.UpdateStatus(ctx, id, "stop")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" stopped\n", name)
	return nil
}

func (u *instanceUseCase) ForceStop(ctx context.Context, name string) error {
	id, err := u.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = u.instanceRepository.UpdateStatus(ctx, id, "forcestop")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" stopped\n", name)
	return nil
}

func (u *instanceUseCase) Delete(ctx context.Context, name string) error {
	id, err := u.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = u.instanceRepository.UpdateStatus(ctx, id, "destroy")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" deleted\n", name)
	return nil
}

func (u *instanceUseCase) getIDFromName(ctx context.Context, name string) (int, error) {
	instances, err := u.instanceRepository.List(ctx)
	if err != nil {
		return 0, err
	}
	for _, element := range instances {
		if element.Name == name {
			return element.ID, nil
		}
	}
	return 0, fmt.Errorf("instance \"%s\" not found", name)
}
