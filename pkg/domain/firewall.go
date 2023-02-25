package domain

import (
	"time"
)

type Firewall struct {
	ID        int
	Name      string    `yaml:"name" print:"NAME"`
	Inbound   []Rule    `yaml:"inbound"`
	Outbound  []Rule    `yaml:"outbound"`
	Instances []string  `yaml:"instances"`
	CreatedAt time.Time `print:"AGE"`
	UpdatedAt time.Time
}

type Rule struct {
	Type     string `yaml:"type" json:"type" print:"TYPE"`
	Protocol string `yaml:"protocol" json:"protocol" print:"PROTOCOL"`
	Port     string `yaml:"port" json:"port" print:"PORT"`
	Source   string `yaml:"source" json:"source" print:"SOURCE"`
}
