package database

import (
	"context"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"github.com/tecwagner/frete_rapido_api/internal/infra/config"
	"gorm.io/gorm"
)

type CarrierDB struct {
	DB *gorm.DB
}

func NewCarrierDB(db *gorm.DB) *CarrierDB {
	return &CarrierDB{
		DB: db,
	}
}

func (c *CarrierDB) Save(ctx context.Context, carriers []entities.Carrier, quoteResponseID uint) error {

	tx := c.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		return config.WrapError(tx.Error, "failed to begin transaction")
	}

	for i := range carriers {
		carriers[i].QuoteResponseID = quoteResponseID

		if err := tx.Create(&carriers[i]).Error; err != nil {
			tx.Rollback()
			return config.WrapError(err, "failed to save carrier")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return config.WrapError(err, "failed to commit transaction")
	}

	return nil
}
