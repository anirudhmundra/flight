package main

import (
	"sahaj/flight/config"
	"sahaj/flight/reader"
	"sahaj/flight/service/flights"
	"sahaj/flight/writer"

	"github.com/sirupsen/logrus"
)

func main() {

	cfg := config.NewConfig()
	csvReader := reader.NewCSVReader(cfg.InputFilePath)
	csvWriter := writer.NewCSVWriter(cfg.ProcessedTicketsFilePath, cfg.FailedTicketsFilePath)
	updater := flights.NewUpadter(csvReader, csvWriter, flights.NewCustomValidtor())
	if err := updater.Update(); err != nil {
		logrus.Error(err)
	}
}
