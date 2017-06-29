package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codeaustin/carbon/utils/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	dbConnection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.Config.DB_User, config.Config.DB_Password, config.Config.DB_Name)

	var err error

	DB, err = sql.Open("postgres", dbConnection)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀  DB Connected")
}
