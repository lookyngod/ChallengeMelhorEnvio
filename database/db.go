package database

import (
	"melhorenvio/pkg/domains/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to database

func ConnectDB() *gorm.DB {
	dsn := "root:Melhor1#@/melhorenvio"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Log{})
	if err != nil {
		panic(err)
	}

	db = db.Begin()

	return db

}
