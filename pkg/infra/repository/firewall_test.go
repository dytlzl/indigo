package repository

import (
	"context"
	"testing"

	_ "embed"

	"github.com/dytlzl/indigo/pkg/domain"
	mock_api "github.com/dytlzl/indigo/pkg/infra/api/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/getFirewall.json
var getFireWallJSON []byte

func Test_apiFirewallRepository_Get(t *testing.T) {
	t.Run("return correct value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		client := mock_api.NewMockClient(ctrl)
		client.EXPECT().Get(context.Background(), "/nw/gettemplate/1").Return(getFireWallJSON, nil)
		a := &apiFirewallRepository{
			Client: client,
		}
		want := domain.Firewall{
			ID:   1,
			Name: "default",
			Inbound: []domain.Rule{
				{Type: "HTTPS", Protocol: "TCP", Port: "443", Source: "0.0.0.0/0"},
				{Type: "HTTP", Protocol: "TCP", Port: "80", Source: "0.0.0.0/0"},
				{Type: "Custom", Protocol: "TCP", Port: "22", Source: "192.168.100.100/32"},
			},
			Outbound: []domain.Rule{
				{Type: "Custom", Protocol: "ICMP", Port: "", Source: "192.168.100.100/32"},
			},
		}
		wantErr := false
		got, err := a.Get(context.Background(), 1)
		if (err != nil) != wantErr {
			t.Errorf("apiFirewallRepository.Get() error = %v, wantErr %v", err, wantErr)
			return
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("apiFirewallRepository.Get(): (-got +want)\n%s", diff)
		}
	})
}
