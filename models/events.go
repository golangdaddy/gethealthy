package models

// NewEvent constructs the event object
func (group *Group) NewEvent(name, description string, start int64) *Event {
	return &Event{
		Meta:        NewInternals(),
		Name:        name,
		Description: description,
		Time:        start,
	}
}

type Event struct {
	Meta        Internals
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
