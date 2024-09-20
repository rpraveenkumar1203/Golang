package models

import (
	"time"

	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

// need to store event

var events = []Event{}

func (e Event) Save() error {

	insertQuery := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)

	`
	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	if err != nil {
		return err

	}
	return err

}

func GetAllEvents() ([]Event, error) {

	getQuery := "SELECT * FROM events"

	eventRows, err := db.DB.Query(getQuery)

	if err != nil {
		return nil, err
	}

	var events []Event

	for eventRows.Next() {
		var event Event

		err := eventRows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}

	return events, nil
}
