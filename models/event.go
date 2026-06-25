package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// later add to db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
