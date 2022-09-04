package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"melhorenvio/database"
	"melhorenvio/pkg/domains/model"
	"melhorenvio/pkg/domains/repository"
	"melhorenvio/pkg/domains/services"
	"os"
	"sync"

	"gorm.io/gorm"
)

func main() {
	db := database.ConnectDB()
	scanner := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}
	var logs []model.Log

	for scanner.Scan() {
		var log model.Log

		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			panic(err)
		}
		logs = append(logs, log)

		fmt.Println("debug bytes", scanner.Bytes())

		wg.Add(1)
		go func(db *gorm.DB, log model.Log) {

			defer wg.Done()

			err := repository.SaveLogInDB(db, log)
			if err != nil {
				panic(err)
			}

		}(db, log)

	}
	fmt.Println("Log saved")

	wg.Add(1)

	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateAverageReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average report generated")

	}(logs)

}
