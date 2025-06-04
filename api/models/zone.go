package models

type Zone struct {
	ID       int64   `json:"id" title:"id" description:"Unique identifier for the zone"`
	Capacity int64   `json:"capacity" title:"capacity" description:"Maximum capacity of the zone"`
	Price    float64 `json:"price" title:"price" description:"Price of the zone in the plan, minimum:0, exclusiveMinimum:true"`
	Name     string  `json:"name" title:"name" description:"Name of the zone, minLength:1"`
	Numbered bool    `json:"numbered" title:"numbered" description:"Indicates if the zone is numbered" enum:"true,false"`
	PlanID   int64   `json:"-"`
	Plan     *Plan   `pg:"rel:belongs-to" json:"-" title:"plan" description:"Plan associated with this zone"`
}
