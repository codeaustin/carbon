package models

import (
	"time"

	"github.com/codeaustin/carbon/utils/db"
)

//Category  used to provide different categories for events
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"string"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//CreateCategoryTable returns the table information to create table in DB
func CreateCategoryTable() db.TableInfo {
	categoryTableInfo := db.TableInfo{
		Name: "category",
		CreationSQL: `
			CREATE TABLE category (
				id SERIAL PRIMARY KEY,
				name VARCHAR NOT NULL,
				created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
				updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
			);`,
	}

	return categoryTableInfo
}
