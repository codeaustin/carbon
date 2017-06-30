package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config struct {
	Port        string
	DB_User     string
	DB_Password string
	DB_Name     string
}

func Init() {
	env := os.Getenv("ENV")

	configDir, _ := os.Getwd()
	configDir += "/config/"

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
