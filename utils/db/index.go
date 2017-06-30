package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codeaustin/carbon/utils/config"
	// Dont actually used but required as driver
	_ "github.com/lib/pq"
)

//DB Globally export DB instance
var DB *sql.DB

//TableInfo used to create all of the necessary tables
type TableInfo struct {
	Name        string
	CreationSQL string
}

//Init connects to the database
func Init() {
	dbConnection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.Config.DB_User, config.Config.DB_Password, config.Config.DB_Name)

	var err error

	DB, err = sql.Open("postgres", dbConnection)

	if err != nil {
		log.Fatal(err)
	}

	// Ensuring connection to DB
	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀  DB Connected")
}

//CreateTables creates tables in DB from the passed TableInfo
func CreateTables(tables []TableInfo) {
	for i := 0; i < len(tables); i++ {
		name := tables[i].Name
		creationSQL := tables[i].CreationSQL

		_, err := DB.Exec("SELECT * FROM " + name)
		if err != nil {
			_, err = DB.Exec(creationSQL)
			fmt.Println("✓ " + name + " table created")
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
