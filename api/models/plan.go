package models

import (
	"time"
)

type Plan struct {
	ID            int64     `xml:"id"`
	PlanStartDate time.Time `xml:"plan_start_date"`
	PlanEndDate   time.Time `xml:"plan_end_date"`
	SellTo        time.Time `xml:"sell_to"`
	SoldOut       bool      `xml:"sold_out"`
	BasePlanID    int64
	BasePlan      *BasePlan `pg:"rel:belongs-to"`
	Zones         []*Zone   `pg:"rel:has-many" xml:"zones"`
}
