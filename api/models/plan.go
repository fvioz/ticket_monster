package models

import (
	"time"
)

type Plan struct {
	ID            int64     `json:"id" title:"id" description:"Unique identifier for the plan"`
	PlanStartDate time.Time `json:"plan_start_date" title:"plan_start_date" description:"Start date of the plan"`
	PlanEndDate   time.Time `json:"plan_end_date" title:"plan_end_date" description:"End date of the plan"`
	SellTo        time.Time `json:"sell_to" title:"sell_to" description:"Deadline for selling the plan tickets"`
	SoldOut       bool      `json:"sold_out" title:"sold_out" description:"Indicates if the plan is sold out"`
	BasePlanID    int64
	BasePlan      *BasePlan `pg:"rel:belongs-to" json:"base_plan" title:"base_plan" description:"Base plan associated with this plan"`
	Zones         []*Zone   `pg:"rel:has-many" json:"zones" title:"zones" description:"List of zones associated with the plan"`
}
