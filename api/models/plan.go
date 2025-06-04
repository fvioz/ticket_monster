package models

import (
	"time"
)

type Plan struct {
	ID            int64     `json:"id" title:"id" description:"Unique identifier for the plan"`
	PlanStartDate time.Time `json:"planStartDate" title:"planStartDate" description:"Start date of the plan"`
	PlanEndDate   time.Time `json:"planEndDate" title:"planEndDate" description:"End date of the plan"`
	SellTo        time.Time `json:"sellTo" title:"sellTo" description:"Deadline for selling the plan tickets"`
	SoldOut       bool      `json:"soldOut" title:"soldOut" description:"Indicates if the plan is sold out"`
	BasePlanID    int64     `json:"-"`
	BasePlan      *BasePlan `pg:"rel:belongs-to" json:"basePlan" title:"base_plan" description:"Base plan associated with this plan"`
	Zones         []*Zone   `pg:"rel:has-many" json:"zones" title:"zones" description:"List of zones associated with the plan"`
}
