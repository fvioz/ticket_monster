package api

import (
	"encoding/xml"
)

type PlanList struct {
	XMLName xml.Name `xml:"planList"`
	Output  Output   `xml:"output"`
}

type Output struct {
	XMLName   xml.Name   `xml:"output"`
	BasePlans []BasePlan `xml:"base_plan"`
}

type BasePlan struct {
	XMLName            xml.Name `xml:"base_plan"`
	ID                 int64    `xml:"base_plan_id,attr"`
	OrganizerCompanyID int64    `xml:"organizer_company_id,attr,omitempty"`
	SellMode           string   `xml:"sell_mode,attr"`
	Title              string   `xml:"title,attr"`
	Plans              []Plan   `xml:"plan"`
}

type Plan struct {
	XMLName       xml.Name `xml:"plan"`
	ID            int64    `xml:"plan_id,attr"`
	PlanStartDate string   `xml:"plan_start_date,attr"`
	PlanEndDate   string   `xml:"plan_end_date,attr"`
	SellTo        string   `xml:"sell_to,attr"`
	SoldOut       bool     `xml:"sold_out,attr"`
	Zones         []Zone   `xml:"zone"`
}

type Zone struct {
	XMLName  xml.Name `xml:"zone"`
	ID       int64    `xml:"zone_id,attr"`
	Name     string   `xml:"name,attr"`
	Capacity int64    `xml:"capacity,attr"`
	Price    string   `xml:"price,attr"`
	Numbered bool     `xml:"numbered"`
}
