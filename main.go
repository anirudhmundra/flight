package main

import (
	"context"
	"sahaj/flight/config"
	"sahaj/flight/reader"
	"sahaj/flight/service/flights"
	"sahaj/flight/writer"

	"github.com/sirupsen/logrus"
)

func main() {

	cfg := config.NewConfig()
	csvReader := reader.NewCSVReader(cfg.InputFilePath)
	csvWriter := writer.NewCSVWriter(cfg.ValidTicketsFilePath, cfg.InvalidTicketsFilePath)
	updater := flights.NewUpadter(csvReader, csvWriter, flights.NewCustomValidtor())
	if err := updater.Update(context.Background()); err != nil {
		logrus.Error(err)
	}
}
