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

type instanceUsecase struct {
	instanceRepository domain.InstanceRepository
}

func NewInstanceUsecase(i domain.InstanceRepository) domain.InstanceUsecase {
	return &instanceUsecase{
		instanceRepository: i,
	}
}

func (i *instanceUsecase) List(ctx context.Context) error {
	instances, err := i.instanceRepository.List(ctx)
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

func (i *instanceUsecase) Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error {
	return i.instanceRepository.Create(ctx, name, planID, osID, regionID, sshKeyID)
}

func (i *instanceUsecase) Start(ctx context.Context, name string) error {
	id, err := i.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = i.instanceRepository.UpdateStatus(ctx, id, "start")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" started\n", name)
	return nil
}

func (i *instanceUsecase) Stop(ctx context.Context, name string) error {
	id, err := i.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = i.instanceRepository.UpdateStatus(ctx, id, "stop")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" stopped\n", name)
	return nil
}

func (i *instanceUsecase) ForceStop(ctx context.Context, name string) error {
	id, err := i.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = i.instanceRepository.UpdateStatus(ctx, id, "forcestop")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" stopped\n", name)
	return nil
}

func (i *instanceUsecase) Delete(ctx context.Context, name string) error {
	id, err := i.getIDFromName(ctx, name)
	if err != nil {
		return err
	}
	err = i.instanceRepository.UpdateStatus(ctx, id, "destroy")
	if err != nil {
		return err
	}
	fmt.Printf("instance \"%s\" deleted\n", name)
	return nil
}

func (i *instanceUsecase) getIDFromName(ctx context.Context, name string) (int, error) {
	instances, err := i.instanceRepository.List(ctx)
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
