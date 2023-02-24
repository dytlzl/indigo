package usecase

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printer"

	"k8s.io/apimachinery/pkg/util/duration"
)

type InstanceRepository interface {
	List(ctx context.Context) ([]domain.Instance, error)
	Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error
	UpdateStatus(ctx context.Context, id int, status string) error
}

type InstanceUsecase interface {
	List(ctx context.Context) error
	Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error
	Start(ctx context.Context, name string) error
	Stop(ctx context.Context, name string) error
	ForceStop(ctx context.Context, name string) error
	Delete(ctx context.Context, name string) error
}

type instanceUsecase struct {
	instanceRepository InstanceRepository
}

func NewInstanceUsecase(i InstanceRepository) InstanceUsecase {
	return &instanceUsecase{
		instanceRepository: i,
	}
}

func (u *instanceUsecase) List(ctx context.Context) error {
	instances, err := u.instanceRepository.List(ctx)
	if err != nil {
		return err
	}
	sort.Slice(instances, func(i, j int) bool { return instances[i].Name < instances[j].Name })
	printer.PrintTable(
		[]string{"NAME", "STATUS", "AGE", "IP", "OS", "PLAN"},
		instances,
		func(instance domain.Instance) []string {
			return []string{instance.Name, instance.Status, duration.HumanDuration(time.Since(instance.StartedAt)), instance.IP, instance.OSName, instance.PlanName}
		},
		"",
	)
	return nil
}

func (u *instanceUsecase) Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error {
	return u.instanceRepository.Create(ctx, name, planID, osID, regionID, sshKeyID)
}

func (u *instanceUsecase) Start(ctx context.Context, name string) error {
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

func (u *instanceUsecase) Stop(ctx context.Context, name string) error {
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

func (u *instanceUsecase) ForceStop(ctx context.Context, name string) error {
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

func (u *instanceUsecase) Delete(ctx context.Context, name string) error {
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

func (u *instanceUsecase) getIDFromName(ctx context.Context, name string) (int, error) {
	instances, err := u.instanceRepository.List(ctx)
	if err != nil {
		return 0, err
	}
	id := func() *int {
		for _, element := range instances {
			if element.Name == name {
				return &element.ID
			}
		}
		return nil
	}()
	if id == nil {
		return 0, fmt.Errorf("instance \"%s\" not found\n", name)
	}
	return *id, nil
}
