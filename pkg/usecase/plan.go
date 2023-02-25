package usecase

import (
	"context"
	"sort"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printutil"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type PlanRepository interface {
	List(ctx context.Context) ([]domain.Plan, error)
}

type PlanUseCase interface {
	List(ctx context.Context) error
}

type planUseCase struct {
	planRepository PlanRepository
}

func NewPlanUseCase(r PlanRepository) PlanUseCase {
	return &planUseCase{
		planRepository: r,
	}
}

func (u *planUseCase) List(ctx context.Context) error {
	plans, err := u.planRepository.List(ctx)
	if err != nil {
		return err
	}
	sort.Slice(plans, func(i, j int) bool { return plans[i].ID < plans[j].ID })
	printutil.PrintTable(plans)
	return nil
}
