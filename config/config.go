package config

import (
	"encoding/json"
	"io/ioutil"
)

type FlightsConfig struct {
	InputFilePath          string `json:"inputFilePath"`
	ValidTicketsFilePath   string `json:"validTicketsFilePath"`
	InvalidTicketsFilePath string `json:"invalidTicketsFilePath"`
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
