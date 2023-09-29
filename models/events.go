package models

import (
	"log"

	"github.com/google/uuid"
)

// NewEvent constructs the event object
func (group *Group) NewEvent(name, description string, start int64, options EventOptions) (event *Event) {
	event = &Event{
		Meta:        NewInternals(),
		Options:     options,
		Group:       group.ID,
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Time:        start,
	}
	log.Println(*event)
	return
}

type Event struct {
	Meta        Internals
	Options     EventOptions
	Group       string  `json:"group" firestore:"group"`
	ID          string  `json:"id" firestore:"id"`
	Name        string  `json:"name" firestore:"name"`
	Description string  `json:"description" firestore:"description"`
	Location    string  `json:"location" firestore:"location"`
	Lat         float64 `json:"lat" firestore:"lat"`
	Lng         float64 `json:"lng" firestore:"lng"`
	Time        int64   `json:"time" firestore:"time"`
	// hours
	Duration float64 `json:"duration" firestore:"duration"`
}

type EventOptions struct {
	// is visible to general search
	Public bool
	// costs no money
	Free bool
	// for accessibility
	DisabledAccess bool
}