package database

import (
	"melhorenvio/pkg/domains/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:Melhor1#@/melhorenvio"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Log{})
	if err != nil {
		panic(err)
	}

	return db

}
