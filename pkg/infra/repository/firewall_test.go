package repository

import (
	"context"
	_ "embed"
	"fmt"
	"testing"
	"time"

	"github.com/dytlzl/indigo/pkg/domain"
	mock_api "github.com/dytlzl/indigo/pkg/infra/api/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/getFirewall.json
var getFireWallJSON []byte

//go:embed testdata/listFirewall.json
var listFireWallJSON []byte

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

func Test_apiFirewallRepository_List(t *testing.T) {
	t.Run("return correct value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		client := mock_api.NewMockClient(ctrl)
		client.EXPECT().Get(NewContextMatcher(3*time.Second, 100*time.Millisecond), "/nw/getfirewalllist").Return(listFireWallJSON, nil)
		a := &apiFirewallRepository{
			Client: client,
		}
		want := []domain.Firewall{
			{
				ID:        1,
				Name:      "default",
				CreatedAt: func() time.Time { createdAt, _ := time.Parse(time.RFC3339, "2022-11-10T12:00:00Z"); return createdAt }(),
			},
		}
		wantErr := false
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		got, err := a.List(ctx)
		if (err != nil) != wantErr {
			t.Errorf("apiFirewallRepository.Get() error = %v, wantErr %v", err, wantErr)
			return
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("apiFirewallRepository.Get(): (-got +want)\n%s", diff)
		}
	})
}

type contextMatcher struct {
	gomock.Matcher
	timeout   time.Duration
	tolerance time.Duration
}

func NewContextMatcher(timeout time.Duration, tolerance time.Duration) gomock.Matcher {
	return contextMatcher{
		timeout:   timeout,
		tolerance: tolerance,
	}
}

func (m contextMatcher) Matches(x interface{}) bool {
	deadline, _ := x.(context.Context).Deadline()
	return time.Until(deadline)-m.timeout < m.tolerance
}

func (m contextMatcher) String() string {
	return fmt.Sprintf("context.Context with timeout %v ±%v", m.timeout, m.tolerance)
}
