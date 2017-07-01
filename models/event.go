package models

import (
	"net/http"
	"time"

	"github.com/codeaustin/carbon/utils/db"
)

type Event struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetEvents() ([]*Event, Tx) {
	var events []*Event

	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, Tx{err.Error(), false, http.StatusInternalServerError}
	}

	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Lat,
			&e.Lon, &e.CreatedAt, &e.UpdatedAt)

		if err != nil {
			return nil, Tx{err.Error(), false, http.StatusInternalServerError}
		}

		events = append(events, &e)
	}

	return events, Tx{"", true, http.StatusOK}
}

func CreateEvent(event *Event) Tx {
	// TODO: validation

	err := db.DB.QueryRow(`INSERT INTO events(title, start_time, end_time, lat, lon) VALUES
		($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`, event.Title, event.StartTime,
		event.EndTime, event.Lat, event.Lon).Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return Tx{err.Error(), false, http.StatusInternalServerError}
	}

	return Tx{"Event successfully created", true, http.StatusCreated}
}

//CreateEventTable returns the table information to create table in DB
func CreateEventTable() db.TableInfo {
	eventTableInfo := db.TableInfo{
		Name: "events",
		CreationSQL: `
			CREATE TABLE events (
				id SERIAL PRIMARY KEY,
				title VARCHAR NOT NULL,
				start_time TIMESTAMP WITHOUT TIME ZONE,
				end_time TIMESTAMP WITHOUT TIME ZONE,
				lat DOUBLE PRECISION NOT NULL,
				lon DOUBLE PRECISION NOT NULL,
				created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
				updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
			);`,
	}

	return eventTableInfo
}
