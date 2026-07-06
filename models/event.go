package models

import (
	"fmt"
	"net/http"
	"time"

	"example.com/rest/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e Event) Save(r *http.Request) (error, uuid.UUID) {

	err := db.Pool.QueryRow(r.Context(),
		"INSERT INTO events (name, description, location, datetime, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		e.Name, e.Description, e.Location, e.DateTime, e.UserId,
	).Scan(&e.Id)
	if err != nil {
		return err, uuid.Nil
	}
	return nil, e.Id
}

func GetAllEvents(r *http.Request) ([]Event, error) {
	rows, err := db.Pool.Query(r.Context(), "SELECT id, name, description, location, datetime, user_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[Event])
	if err != nil {
		return nil, err
	}
	return events, nil
}

func GetEventById(r *http.Request, id string) (Event, error) {
	var event Event
	err := db.Pool.QueryRow(r.Context(), "SELECT id, name, description, location, datetime, user_id FROM events WHERE id=$1", id).Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
func DeleteEventById(r *http.Request, id string) error {
	tag, err := db.Pool.Exec(r.Context(), "DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("event not found")
	}
	return nil
}
