package usecase

import (
	"context"
	"errors"
	"testing"

	mock_usecase "github.com/dytlzl/indigo/pkg/usecase/mock"
	"github.com/golang/mock/gomock"
)

func Test_instanceUseCase_Create(t *testing.T) {
	t.Run("return nil when instance was created successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		r := mock_usecase.NewMockInstanceRepository(ctrl)
		u := &instanceUseCase{
			instanceRepository: r,
		}
		r.EXPECT().Create(context.Background(), "instance01", 1, 1, 1, 1).Return(nil)
		wantErr := false
		if err := u.Create(context.Background(), "instance01", 1, 1, 1, 1); (err != nil) != wantErr {
			t.Errorf("instanceUseCase.Create() error = %v, wantErr %v", err, wantErr)
		}
	})
	t.Run("return err when it failed to create instance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		r := mock_usecase.NewMockInstanceRepository(ctrl)
		u := &instanceUseCase{
			instanceRepository: r,
		}
		r.EXPECT().Create(context.Background(), "instance01", 1, 1, 1, 1).Return(errors.New("failed"))
		wantErr := true
		if err := u.Create(context.Background(), "instance01", 1, 1, 1, 1); (err != nil) != wantErr {
			t.Errorf("instanceUseCase.Create() error = %v, wantErr %v", err, wantErr)
		}
	})
}
