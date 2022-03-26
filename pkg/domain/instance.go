package domain

import (
	"context"
	"time"
)

type Instance struct {
	ID        int
	Name      string
	IP        string
	OSName    string
	Status    string
	StartedAt time.Time
	PlanName  string
}

type InstanceRepository interface {
	List(ctx context.Context) ([]Instance, error)
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
