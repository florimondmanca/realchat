package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Configuration stores the server configuration parameters
type Configuration struct {
	Port        int    `json:"port"`
	Host        string `json:"host"`
	LogFilePath string `json:"logFilePath"`
}

func load() Configuration {
	configuration := Configuration{}
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		log.Panic("Error loading configuration:", err.Error())
		panic("Error loading configuration")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Panic("Error parsing configuration file:", err.Error())
		panic("Error loading configuration")
	}

	return configuration
}

func setupLog(configuration Configuration) {
	logFile, err := os.Create(configuration.LogFilePath)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
}

// Load loads config, sets up logging and returns config
func Load() Configuration {
	configuration := load()
	setupLog(configuration)
	return configuration
}
