package domain

import (
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
