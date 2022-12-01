package config

import (
	"encoding/json"
	"io/ioutil"
)

type FlightsConfig struct {
	InputFilePath            string `json:"inputFilePath"`
	ProcessedTicketsFilePath string `json:"processedTicketsFilePath"`
	FailedTicketsFilePath    string `json:"failedTicketsFilePath"`
}

func NewConfig() FlightsConfig {
	config, err := ioutil.ReadFile(`config\config.json`)
	if err != nil {
		panic(err)
	}

	var flightsConfig FlightsConfig
	if err := json.Unmarshal(config, &flightsConfig); err != nil {
		panic(err)
	}

	return flightsConfig
}
