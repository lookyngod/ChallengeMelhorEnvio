package services

import (
	"melhorenvio/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SaveLog(db *gorm.DB, log models.Log) error {
	log.ID = uuid.New()
	result := db.Create(log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
