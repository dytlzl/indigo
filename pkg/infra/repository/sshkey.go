package repository

import (
	"context"
	"encoding/json"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/usecase"
)

type apiSSHKeyRepository struct {
	Client api.Client
}

type sshKeyListResponse struct {
	SSHKeys []domain.SSHKey `json:"sshkeys"`
}

func NewAPISSHKeyRepository(client api.Client) usecase.SSHKeyRepository {
	return &apiSSHKeyRepository{Client: client}
}

func (a *apiSSHKeyRepository) List(ctx context.Context) ([]domain.SSHKey, error) {
	bytes, err := a.Client.Get(ctx, "/vm/sshkey")
	if err != nil {
		return nil, err
	}

	sshKeyList := sshKeyListResponse{}
	err = json.Unmarshal(bytes, &sshKeyList)
	if err != nil {
		return nil, err
	}
	return sshKeyList.SSHKeys, nil
}
