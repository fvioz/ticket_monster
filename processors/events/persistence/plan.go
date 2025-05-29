package persistence

import (
	"time"
)

type Plan struct {
	ID            int64
	PlanStartDate time.Time
	PlanEndDate   time.Time
	SellTo        time.Time
	SoldOut       bool
	Zones         []Zone
}
