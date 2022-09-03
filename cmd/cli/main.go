package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"melhorenvio/database"
	"melhorenvio/pkg/domains/model"
	"melhorenvio/pkg/domains/repository"
	"os"
	"sync"

	"gorm.io/gorm"
)

func main() {
	db := database.ConnectDB()

	scanner := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}

	for scanner.Scan() {
		var log model.Log

		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			panic(err)
		}

		fmt.Println("debug bytes", scanner.Bytes())

		wg.Add(1)
		go func(db *gorm.DB, log model.Log) {

			fmt.Println("go func rodando")
			defer wg.Done()

			err := repository.SaveLog(db, log)
			if err != nil {
				panic(err)
			}
			fmt.Println("Log saved")

		}(db, log)

		fmt.Println("chegou aqui")

	}
}
