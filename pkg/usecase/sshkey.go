package usecase

import (
	"context"
	"sort"
	"strconv"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printer"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type SSHKeyRepository interface {
	List(ctx context.Context) ([]domain.SSHKey, error)
}

type SSHKeyUseCase interface {
	List(ctx context.Context) error
}

type sshKeyUseCase struct {
	sshKeyRepository SSHKeyRepository
}

func NewSSHKeyUseCase(r SSHKeyRepository) SSHKeyUseCase {
	return &sshKeyUseCase{
		sshKeyRepository: r,
	}
}

func (u *sshKeyUseCase) List(ctx context.Context) error {
	sshKeys, err := u.sshKeyRepository.List(ctx)
	if err != nil {
		return err
	}
	sort.Slice(sshKeys, func(i, j int) bool { return sshKeys[i].Name < sshKeys[j].Name })
	printer.PrintTable(
		[]string{"NAME", "ID", "STATUS"},
		sshKeys,
		func(sshKey domain.SSHKey) []string {
			return []string{sshKey.Name, strconv.Itoa(sshKey.ID), sshKey.Status}
		},
		"",
	)
	return nil
}
