package models

type BasePlan struct {
	ID       int64  `json:"id" title:"id" description:"Unique identifier for the base plan"`
	SellMode string `json:"sell_mode" title:"sell_mode" description:"Mode of selling the plan,enum=online,enum=offline"`
	Title    string `json:"title" title:"title" description:"Title of the base plan"`
	Plans    []Plan `pg:"rel:has-many" json:"plans" title:"plans" description:"List of plans associated with the base plan"`
}
