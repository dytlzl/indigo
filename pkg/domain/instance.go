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
