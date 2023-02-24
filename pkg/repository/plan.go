package repository

import (
	"context"
	_ "embed"
	"encoding/json"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/usecase"
)

func NewJSONPlanRepository() usecase.PlanRepository {
	return &jsonPlanRepository{}
}

type jsonPlanRepository struct {
}

//go:embed data/plans.json
var plansJsonBytes []byte

func (r *jsonPlanRepository) List(context.Context) ([]domain.Plan, error) {
	plans := make([]domain.Plan, 0)
	err := json.Unmarshal(plansJsonBytes, &plans)
	return plans, err
}
