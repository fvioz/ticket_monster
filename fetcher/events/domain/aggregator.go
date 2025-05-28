package domain

import (
	"fetcher/events/api"
	"fetcher/libs"
	"strconv"
	"time"
)

type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

// Convert the XML BasePlan slice to a slice of BasePlan
func (a *Aggregator) FromApiToDomain(apiBasePlan []api.BasePlan) []BasePlan {
	var basePlans []BasePlan

	for _, xmlBasePlan := range apiBasePlan {
		newBasePlan := a.FromApiBasePlanToDomain(xmlBasePlan)
		basePlans = append(basePlans, newBasePlan)
	}

	return basePlans
}

func (r *Aggregator) FromApiBasePlanToDomain(basePlan api.BasePlan) BasePlan {
	return BasePlan{
		ID:       basePlan.ID,
		SellMode: basePlan.SellMode,
		Title:    basePlan.Title,
		Plans:    r.FromApiPlansToDomain(basePlan.Plans),
	}
}

func (r *Aggregator) FromApiPlansToDomain(xmlPlans []api.Plan) []Plan {
	var plans []Plan

	for _, xmlPlan := range xmlPlans {
		newPlan := r.FromApiPlanToDomain(xmlPlan)
		plans = append(plans, newPlan)
	}

	return plans
}

func (r *Aggregator) FromApiPlanToDomain(plan api.Plan) Plan {
	return Plan{
		ID:            plan.ID,
		PlanStartDate: r.FromTimeStringToTime(plan.PlanStartDate),
		PlanEndDate:   r.FromTimeStringToTime(plan.PlanEndDate),
		SellTo:        r.FromTimeStringToTime(plan.SellTo),
		SoldOut:       plan.SoldOut,
		Zones:         r.FromApiZonesToDomain(plan.Zones),
	}
}

func (r *Aggregator) FromApiZonesToDomain(xmlZones []api.Zone) []Zone {
	var zones []Zone

	for _, xmlZone := range xmlZones {
		newZone := r.FromApiZoneToDomain(xmlZone)
		zones = append(zones, newZone)
	}

	return zones
}

func (r *Aggregator) FromApiZoneToDomain(zone api.Zone) Zone {
	price, err := strconv.ParseFloat(zone.Price, 64)
	if err != nil {
		price = 0.0 // Default to 0.0 if parsing fails
	}

	return Zone{
		ID:       zone.ID,
		Capacity: zone.Capacity,
		Price:    price,
		Name:     zone.Name,
		Numbered: zone.Numbered,
	}
}

func (r *Aggregator) FromTimeStringToTime(t string) time.Time {
	logger := libs.LoggerInstance()

	layout := "2006-01-02T15:04:05"
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		logger.Error("Error parsing time:", err)
		return time.Time{} // Return zero time if parsing fails
	}

	return parsedTime
}
