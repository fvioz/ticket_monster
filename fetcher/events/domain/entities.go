package domain

import (
	"time"
)

type BasePlan struct {
	ID       int64
	SellMode string
	Title    string
	Plans    []Plan
}

type Plan struct {
	ID            int64
	PlanStartDate time.Time
	PlanEndDate   time.Time
	SellTo        time.Time
	SoldOut       bool
	Zones         []Zone
}

type Zone struct {
	ID       int64
	Capacity int64
	Price    float64
	Name     string
	Numbered bool
}
