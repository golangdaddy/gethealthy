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
	ID          string  `json:"id" firestore:"id"`
	Name        string  `json:"name" firestore:"name"`
	Description string  `json:"description" firestore:"description"`
	Location    string  `json:"location" firestore:"location"`
	Lat         string  `json:"lat" firestore:"lat"`
	Lng         string  `json:"lng" firestore:"lng"`
	Time        int64   `json:"time" firestore:"time"`
	Duration    float64 `json:"duration" firestore:"duration"`
}
