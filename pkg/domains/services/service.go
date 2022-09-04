package services

import (
	"encoding/csv"
	"fmt"
	"melhorenvio/pkg/domains/model"
	"os"
)

func GenerateAverageReport(logs []model.Log) error {

	records := [][]string{
		{"service", "request_time", "proxy_time", "gateway_time"},
	}

	recordMap := make(map[string]model.AverageRecord, 0)

	for _, log := range logs {
		record := recordMap[log.Service.Name]

		record.Quantity += 1
		record.RequestTime += log.Latencies.Request
		record.ProxyTime += log.Latencies.Proxy
		record.GatewayTime += log.Latencies.Kong

		recordMap[log.Service.Name] = record

	}

	fmt.Println("CHEGOU aqui")

	for service, record := range recordMap {
		records = append(records, []string{
			fmt.Sprintf("%v", service),
			fmt.Sprintf("%v", record.RequestTime/record.Quantity),
			fmt.Sprintf("%v", record.ProxyTime/record.Quantity),
			fmt.Sprintf("%v", record.GatewayTime/record.Quantity),
		})
	}

	fmt.Println("CHEGOU AQUI2")

	file, err := os.Create("report_time_average.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("CHEGOU AQUI3")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}
