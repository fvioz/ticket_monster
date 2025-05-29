package domain

import (
	"time"
)

type BasePlan struct {
	ID       int64
	SellMode string
	Title    string
	Plans    []*Plan `pg:"rel:has-many,join_fk:base_plan_id"`
}

type Plan struct {
	ID            int64
	PlanStartDate time.Time
	PlanEndDate   time.Time
	SellTo        time.Time
	SoldOut       bool
	BasePlan      *BasePlan `pg:"rel:has-one"`
	Zones         []*Zone   `pg:"rel:has-many"`
}

type Zone struct {
	ID       int64
	Capacity int64
	Price    float64
	Name     string
	Numbered bool
	Plan     *Plan `pg:"rel:has-one"`
}
