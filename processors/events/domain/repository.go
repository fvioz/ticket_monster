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

	bp := &BasePlan{ID: basePlan.ID}
	err := db.Model(bp).WherePK().Select()
	if err != nil {
		tx, err := db.Begin()
		if err != nil {
			logger.Error("Failed to begin transaction", zap.Error(err))
			return err
		}
		defer tx.Close()

		_, err = db.Model(basePlan).Insert()
		if err != nil {
			_ = tx.Rollback()
			logger.Error("Failed to insert base plan, plans and zones", zap.Error(err))
			return err
		}
	}

	return nil
}
