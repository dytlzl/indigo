package domain

import (
	"time"
)

type Instance struct {
	ID        int
	Name      string    `print:"0,NAME"`
	IP        string    `print:"3,IP"`
	OSName    string    `print:"4,OS"`
	Status    string    `print:"1,STATUS"`
	StartedAt time.Time `print:"2,AGE"`
	PlanName  string    `print:"5,PLAN"`
}

type OS struct {
	ID   int    `print:"1,ID"`
	Name string `print:"0,NAME"`
}

type Plan struct {
	ID      int    `print:",ID"`
	Code    string `print:",CODE"`
	VCPU    int    `print:",VCPU"`
	RAM     int    `print:",RAM"`
	SSD     int    `print:",SSD"`
	IPType  string `print:",IP TYPE"`
	Network string `print:",NETWORK"`
}

type SSHKey struct {
	ID     int    `print:"1,ID"`
	Name   string `print:"0,NAME"`
	Status string `print:"2,STATUS"`
}
