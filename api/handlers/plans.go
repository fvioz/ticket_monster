package handlers

import (
	"api/libs"
	"api/models"
	"time"
)

type PlanHandler struct{}

func NewPlansHandler() *PlanHandler {
	return &PlanHandler{}
}

func (p *PlanHandler) GetPlansV1(starts_at time.Time, ends_at time.Time) ([]models.Plan, error) {
	db := libs.DBInstance()

	var plans []models.Plan

	db.Model(&plans).
		Relation("BasePlan").
		Relation("Zones").
		Where(
			"base_plan.sell_mode >= ? AND plan.plan_start_date >= ? AND plan.plan_end_date <= ?",
			"online",
			starts_at,
			ends_at,
		).
		Select()

	return plans, nil
}
