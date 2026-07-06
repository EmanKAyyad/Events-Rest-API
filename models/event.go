package models

import (
	"net/http"
	"time"

	"example.com/rest/db"
	"github.com/google/uuid"
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

// func GetAllEvents() []Event {
// return events
// }
