package usecase

import (
	"context"
	"sort"
	"strconv"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printutil"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type OSRepository interface {
	List(ctx context.Context) ([]domain.OS, error)
}

type OSUseCase interface {
	List(ctx context.Context) error
}

type osUseCase struct {
	osRepository OSRepository
}

func NewOSUseCase(r OSRepository) OSUseCase {
	return &osUseCase{
		osRepository: r,
	}
}

func (u *osUseCase) List(ctx context.Context) error {
	oses, err := u.osRepository.List(ctx)
	if err != nil {
		return err
	}
	sort.Slice(oses, func(i, j int) bool { return oses[i].Name < oses[j].Name })
	printutil.PrintTable(
		[]string{"NAME", "ID"},
		oses,
		func(os domain.OS) []string {
			return []string{os.Name, strconv.Itoa(os.ID)}
		},
		"",
	)
	return nil
}
