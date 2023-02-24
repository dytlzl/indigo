package domain

import (
	"time"
)

type Instance struct {
	ID        int
	Name      string
	IP        string
	OSName    string
	Status    string
	StartedAt time.Time
	PlanName  string
}

type OS struct {
	ID   int
	Name string
}

type Plan struct {
	ID      int
	Code    string
	IPType  string
	VCPU    int
	RAM     int
	SSD     int
	Network string
}

type SSHKey struct {
	ID     int
	Name   string
	Status string
}
