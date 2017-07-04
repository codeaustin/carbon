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

func GetEvents(limit, offset uint64) ([]*Event, Tx) {
	events := make([]*Event, 0)

	rows, err := db.DB.Query("SELECT * FROM events ORDER BY created_at DESC LIMIT $1 OFFSET $2;",
		limit, offset)

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

func GetEvent(id string) (*Event, Tx) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id=$1;", id)

	var e Event
	err := row.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Lat,
		&e.Lon, &e.CreatedAt, &e.UpdatedAt)

	if err != nil {
		return nil, Tx{err.Error(), false, http.StatusInternalServerError}
	}
	return &e, Tx{"", true, http.StatusOK}
}

func DeleteEvent(id string) Tx {
	row := db.DB.QueryRow("DELETE FROM events WHERE id=$1 RETURNING *;", id)

	var e Event
	err := row.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Lat,
		&e.Lon, &e.CreatedAt, &e.UpdatedAt)

	if err != nil {
		return Tx{err.Error(), false, http.StatusInternalServerError}
	}
	return Tx{"Event successfully deleted", true, http.StatusOK}
}

func CreateEvent(event *Event) Tx {
	// TODO: validation

	err := db.DB.QueryRow(`INSERT INTO events(title, start_time, end_time, lat, lon) VALUES
		($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at;`, event.Title, event.StartTime,
		event.EndTime, event.Lat, event.Lon).Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return Tx{err.Error(), false, http.StatusInternalServerError}
	}

	return Tx{"Event successfully created", true, http.StatusCreated}
}

func UpdateEvent(id string, fieldsMap map[string]interface{}) Tx {
	for key, value := range fieldsMap {

		query := fmt.Sprintf("UPDATE events SET %s='%v' WHERE id=%s RETURNING *", key, value, id)
		row := db.DB.QueryRow(query)

		var e Event
		err := row.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Lat,
			&e.Lon, &e.CreatedAt, &e.UpdatedAt)

		if err != nil {
			return Tx{err.Error(), false, http.StatusInternalServerError}
		}
	}

	return Tx{"Event successfully updated", true, http.StatusOK}
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
