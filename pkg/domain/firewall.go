package domain

import (
	"context"
	"time"
)

type Firewall struct {
	ID        int
	Name      string   `yaml:"name"`
	Inbound   []Rule   `yaml:"inbound"`
	Outbound  []Rule   `yaml:"outbound"`
	Instances []string `yaml:"instances"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Rule struct {
	Type     string `yaml:"type" json:"type"`
	Protocol string `yaml:"protocol" json:"protocol"`
	Port     string `yaml:"port" json:"port"`
	Source   string `yaml:"source" json:"source"`
}

type FirewallRepository interface {
	List(ctx context.Context) ([]Firewall, error)
	Get(ctx context.Context, id int) (*Firewall, error)
	Create(ctx context.Context, fw *Firewall) error
	Update(ctx context.Context, fw *Firewall) error
	Delete(ctx context.Context, id int) error
}

type FirewallUsecase interface {
	List(ctx context.Context) error
	Get(ctx context.Context, target string) error
	Apply(ctx context.Context, fileBody []byte) error
	Delete(ctx context.Context, target string) error
}
