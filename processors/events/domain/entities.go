package domain

import (
	"time"
)

type BasePlan struct {
	ID       int64
	SellMode string
	Title    string
	Plans    []*Plan `pg:"rel:has-many"`
}

type Plan struct {
	ID            int64
	PlanStartDate time.Time
	PlanEndDate   time.Time
	SellTo        time.Time
	SoldOut       bool
	BasePlan      *BasePlan `pg:"rel:belongs-to"`
	Zones         []*Zone   `pg:"rel:has-many"`
}

type Zone struct {
	ID       int64
	Capacity int64
	Price    float64
	Name     string
	Numbered bool
	Plan     *Plan `pg:"rel:belongs-to"`
}
