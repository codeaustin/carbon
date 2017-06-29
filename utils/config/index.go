package config

import (
	"os"
	"encoding/json"
	"log"
)


var Config struct {
	Port string
	DB_User string
	DB_Password string
	DB_Name string
}


func Load(configPath string) {
	env := os.Getenv("ENV")

	var configFile string
	if env == "PROD" {
		configFile = configPath + "/config.prod.json"
	} else if env == "TEST" {
		configFile = configPath + "/config.test.json"
	} else {
		configFile = configPath + "/config.dev.json"
	}

	file, err := os.Open(configFile); if err != nil { log.Fatal(err) }

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)
	
	if err != nil { 
		log.Fatal("There was a problem parsing the config json")
	}
}