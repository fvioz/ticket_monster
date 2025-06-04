package models

type BasePlan struct {
	ID           int64  `json:"id" title:"id" description:"Unique identifier for the base plan"`
	SellMode     string `json:"sellMode" title:"sell_mode" description:"Mode of selling the plan,enum=online,enum=offline"`
	QueueLimit   int64  `json:"queueLimit" title:"queue_limit" description:"Maximum number of clients allowed in the queue"`
	QueueEnabled bool   `json:"queueEnabled" title:"queue_enabled" description:"Indicates if the queue is enabled for this plan"`
	Title        string `json:"title" title:"title" description:"Title of the base plan"`
	Plans        []Plan `pg:"rel:has-many" json:"-" title:"plans" description:"List of plans associated with the base plan"`
}
