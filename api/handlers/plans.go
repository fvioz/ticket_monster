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

type Plans struct {
	Plans []models.Plan `json:"plans" title:"plans" description:"List of plans available for the specified period"`
}

func (p *PlanHandler) GetPlansV1(starts_at time.Time, ends_at time.Time) (Plans, error) {
	db := libs.DBInstance()

	var plans []models.Plan

	db.Model(&plans).
		Relation("BasePlan").
		Relation("Zones").
		Where(
			"base_plan.sell_mode = ? AND plan.plan_start_date >= ? AND plan.plan_end_date <= ?",
			"online",
			starts_at,
			ends_at,
		).
		Select()

	return Plans{Plans: plans}, nil
}
