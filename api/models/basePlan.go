package models

type BasePlan struct {
	ID       int64  `xml:"id"`
	SellMode string `xml:"sell_mode"`
	Title    string `xml:"title"`
	Plans    []Plan `pg:"rel:has-many" xml:"plans"`
}
