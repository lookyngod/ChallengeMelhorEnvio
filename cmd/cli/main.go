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

		fmt.Println("Number of logs: ", len(logs))

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

	wg.Add(3)

	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateAverageTimeServiceReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average report generated")

	}(logs)

	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateRequestPerConsumerReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average per consumer report generated")

	}(logs)

	go func(logs []model.Log) {

		defer wg.Done()

		err := services.GenerateRequestPerServicesReport(logs)
		if err != nil {
			panic(err)
		}
		fmt.Println("Time average per services report generated")

	}(logs)

	wg.Wait()

	db.Commit()

}
