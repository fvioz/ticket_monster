package domain

import (
	"processors/libs"

	"go.uber.org/zap"
)

type Repository struct {
	aggregator *Aggregator
}

func NewRepository() *Repository {
	return &Repository{
		aggregator: NewAggregator(),
	}
}

func (r *Repository) SaveEvent(basePlan BasePlan) error {
	logger := libs.LoggerInstance()
	db := libs.DBInstance()

	logger.Info("Inserting base plan",
		zap.Int64("ID", basePlan.ID),
		zap.String("title", basePlan.Title),
	)

	tx, err := db.Begin()
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return err
	}
	defer tx.Close()

	_, err = tx.Model(&basePlan).
		Where("id = ?id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		_ = tx.Rollback()
		logger.Error("Failed to insert base plan", zap.Error(err))
		return err
	}

	for _, plan := range basePlan.Plans {
		plan.BasePlanID = basePlan.ID
		_, err = tx.Model(plan).
			Where("id = ?id").
			OnConflict("DO NOTHING").
			SelectOrInsert()
		if err != nil {
			_ = tx.Rollback()
			logger.Error("Failed to insert plan", zap.Error(err))
			return err
		}

		for _, zone := range plan.Zones {
			zone.PlanID = plan.ID
			_, err = tx.Model(zone).
				Where("id = ?id").
				OnConflict("DO NOTHING").
				SelectOrInsert()
			if err != nil {
				_ = tx.Rollback()
				logger.Error("Failed to insert zone", zap.Error(err))
				return err
			}
		}

		if err := tx.Commit(); err != nil {
			logger.Error("Failed to insert the event", zap.Error(err))
		}
	}

	return nil
}
