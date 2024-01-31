package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UsersID     int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to a db
	events = append(events, e)
}
