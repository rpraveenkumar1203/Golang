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
	UserID      int64
}

// need to store event

func (e *Event) Save() error {

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

func GetEventbyID(id int64) (*Event, error) {

	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, err

}

func (e Event) UpdateEvent() error {

	query :=
		`
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err

}

func (e Event) DeleteEvent() error {

	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}
