package domain

import (
	"processors/events/persistence"
)

type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

func (a *Aggregator) FromPersistenceToDomain(persistenceBasePlan []persistence.BasePlan) []BasePlan {
	var basePlans []BasePlan

	for _, jsonBasePlan := range persistenceBasePlan {
		newBasePlan := a.FromPersistenceBasePlanToDomain(jsonBasePlan)
		basePlans = append(basePlans, newBasePlan)
	}

	return basePlans
}

func (r *Aggregator) FromPersistenceBasePlanToDomain(persistenceBasePlan persistence.BasePlan) BasePlan {
	return BasePlan{
		ID:       persistenceBasePlan.ID,
		SellMode: persistenceBasePlan.SellMode,
		Title:    persistenceBasePlan.Title,
		Plans:    r.FromPersistencePlansToDomain(persistenceBasePlan.Plans),
	}
}

func (r *Aggregator) FromPersistencePlansToDomain(persistencePlans []persistence.Plan) []*Plan {
	var plans []*Plan

	for _, persistencePlan := range persistencePlans {
		newPlan := r.FromPersistencePlanToDomain(persistencePlan)
		plans = append(plans, newPlan)
	}

	return plans
}

func (r *Aggregator) FromPersistencePlanToDomain(persistenceBasePlanPlan persistence.Plan) *Plan {
	return &Plan{
		ID:            persistenceBasePlanPlan.ID,
		PlanStartDate: persistenceBasePlanPlan.PlanStartDate,
		PlanEndDate:   persistenceBasePlanPlan.PlanEndDate,
		SellTo:        persistenceBasePlanPlan.SellTo,
		SoldOut:       persistenceBasePlanPlan.SoldOut,
		Zones:         r.FromPersistenceZonesToDomain(persistenceBasePlanPlan.Zones),
	}
}

func (r *Aggregator) FromPersistenceZonesToDomain(persistenceBasePlanPlanZones []persistence.Zone) []*Zone {
	var zones []*Zone

	for _, persistenceZone := range persistenceBasePlanPlanZones {
		newZone := r.FromPersistenceZoneToDomain(persistenceZone)
		zones = append(zones, newZone)
	}

	return zones
}

func (r *Aggregator) FromPersistenceZoneToDomain(persistenceZone persistence.Zone) *Zone {
	return &Zone{
		ID:       persistenceZone.ID,
		Capacity: persistenceZone.Capacity,
		Price:    persistenceZone.Price,
		Name:     persistenceZone.Name,
		Numbered: persistenceZone.Numbered,
	}
}
