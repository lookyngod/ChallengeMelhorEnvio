package repository

import (
	"melhorenvio/pkg/domains/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Function to save log in database

func SaveLogInDB(db *gorm.DB, log model.Log) error {
	log.ID = uuid.New()
	result := db.Create(log)
	if result.Error != nil {
		db.Rollback()
		return result.Error
	}

	return nil
}
