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
	// Connect to database
	db := database.ConnectDB()
	// Scan file input
	scanner := bufio.NewScanner(os.Stdin)

	// Wait group to wait all goroutines finish
	wg := sync.WaitGroup{}

	var logs []model.Log

	// Read the file line by line
	for scanner.Scan() {
		var log model.Log

		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			panic(err)
		}

		logs = append(logs, log)

		fmt.Println("Number of logs: ", len(logs))

		wg.Add(1)
		// Save log in database
		go func(db *gorm.DB, log model.Log) {

			defer wg.Done()

			err := repository.SaveLogInDB(db, log)
			if err != nil {
				panic(err)
			}

		}(db, log)

	}
	fmt.Println("Log saved")
	wg.Add(3)
	// Generate report
	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateAverageTimeServicesReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average report generated")

	}(logs)

	// Generate report
	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateRequestPerConsumerReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average per consumer report generated")

	}(logs)

	// Generate report
	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateRequestPerServiceReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average per services report generated")

	}(logs)
	// Wait all goroutines finish
	wg.Wait()
	// Commit transaction
	db.Commit()

}
