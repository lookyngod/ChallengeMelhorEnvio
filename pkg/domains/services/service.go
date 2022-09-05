package services

import (
	"encoding/csv"
	"fmt"
	"melhorenvio/pkg/domains/model"
	"os"
)

func GenerateAverageTimeServiceReport(logs []model.Log) error {

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

	for service, record := range recordMap {
		records = append(records, []string{
			fmt.Sprintf("%v", service),
			fmt.Sprintf("%v", record.RequestTime/record.Quantity),
			fmt.Sprintf("%v", record.ProxyTime/record.Quantity),
			fmt.Sprintf("%v", record.GatewayTime/record.Quantity),
		})
	}

	file, err := os.Create("report_time_average.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

func GenerateRequestPerConsumerReport(logs []model.Log) error {

	records := [][]string{
		{"services", "requests"},
	}

	recordMap := make(map[string]int64, 0)

	for _, log := range logs {

		recordMap[log.Service.Name] += 1

	}

	for services, record := range recordMap {
		records = append(records, []string{
			fmt.Sprintf("%v", services),
			fmt.Sprintf("%v", record),
		})
	}

	file, err := os.Create("report_request_per_services.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

func GenerateRequestPerServicesReport(logs []model.Log) error {

	records := [][]string{
		{"consumer", "requests"},
	}

	recordMap := make(map[string]int64, 0)

	for _, log := range logs {

		recordMap[log.AuthenticatedEntity.ConsumerID.UUID] += 1

	}

	for consumer, record := range recordMap {
		records = append(records, []string{
			fmt.Sprintf("%v", consumer),
			fmt.Sprintf("%v", record),
		})
	}

	file, err := os.Create("report_request_per_consumer.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}
