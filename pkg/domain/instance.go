package domain

import (
	"time"
)

type Instance struct {
	ID        int
	Name      string    `print:"NAME"`
	Status    string    `print:"STATUS"`
	StartedAt time.Time `print:"AGE"`
	IP        string    `print:"IP"`
	OSName    string    `print:"OS"`
	PlanName  string    `print:"PLAN"`
}

type OS struct {
	ID   int    `print:"ID,1"`
	Name string `print:"NAME"`
}

type Plan struct {
	ID      int    `print:"ID"`
	Code    string `print:"CODE"`
	VCPU    int    `print:"VCPU"`
	RAM     int    `print:"RAM"`
	SSD     int    `print:"SSD"`
	IPType  string `print:"IP TYPE"`
	Network string `print:"NETWORK"`
}

type SSHKey struct {
	ID     int    `print:"ID,1"`
	Name   string `print:"NAME"`
	Status string `print:"STATUS,2"`
}
