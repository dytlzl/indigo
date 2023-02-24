package repository

import (
	"context"
	"encoding/json"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/usecase"
)

type apiOSRepository struct {
	Client api.Client
}

type osListResponse struct {
	OSCategory []osCategory `json:"osCategory"`
}

type osCategory struct {
	Name    string `json:"name"`
	OSLists []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"osLists"`
}

func NewAPIOSRepository(client api.Client) usecase.OSRepository {
	return &apiOSRepository{Client: client}
}

func (a *apiOSRepository) List(ctx context.Context) ([]domain.OS, error) {
	bytes, err := a.Client.Get(ctx, "/vm/oslist?instanceTypeId=1")
	if err != nil {
		return nil, err
	}

	osList := osListResponse{}
	err = json.Unmarshal(bytes, &osList)
	if err != nil {
		return nil, err
	}
	oses := make([]domain.OS, 0)
	for _, categories := range osList.OSCategory {
		for _, os := range categories.OSLists {
			oses = append(oses, domain.OS{
				ID:   os.ID,
				Name: os.Name,
			})
		}
	}
	return oses, nil
}
