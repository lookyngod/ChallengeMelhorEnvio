package repository

import (
	"melhorenvio/pkg/domains/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SaveLogInDB(db *gorm.DB, log model.Log) error {
	log.ID = uuid.New()
	result := db.Create(log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
