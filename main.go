package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"melhorenvio/models"
	"os"
	"sync"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Melhor1#@/melhorenvio"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Log{})
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}

	for scanner.Scan() {
		var log models.Log

		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			panic(err)
		}

		fmt.Println("debug bytes", scanner.Bytes())

		wg.Add(1)
		go func(db *gorm.DB, log models.Log) {

			fmt.Println("go func rodando")
			defer wg.Done()

			err := saveLog(db, log)
			if err != nil {
				panic(err)
			}
			fmt.Println("Log saved")

		}(db, log)

		fmt.Println("chegou aqui")

	}
}

func saveLog(db *gorm.DB, log models.Log) error {
	log.ID = uuid.New()
	result := db.Create(log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
