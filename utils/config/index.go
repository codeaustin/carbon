package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config struct {
	Port       string `json:"port"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBPort     string `json:"dbPort"`
	DBHost     string `json:"dbHost"`
}

func init() {
	env := os.Getenv("ENV")

	configDir, _ := os.Getwd()
	configDir += "/config"

	var configFile string
	if env == "PROD" {
		configFile = configDir + "/config.prod.json"
	} else if env == "TEST" {
		configFile = configDir + "/config.test.json"
	} else {
		configFile = configDir + "/config.dev.json"
	}

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatal("There was a problem parsing the config json")
	}
}
