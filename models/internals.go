package models

import "time"

// NewInternals returns a boilerplate internal object
func NewInternals() Internals {

	timestamp := time.Now().UTC().Unix()

	i := Internals{}
	i.Created = timestamp
	i.Modified = timestamp
	return i
}

type Internals struct {
	Moderation struct {
		Blocked  bool
		Approved bool
	}
	Searchable bool
	Created    int64
	Modified   int64
	Stats      struct {
		Followers int64
		Views     int64
		Likes     int64
		Replies   int64
	}
}

func (i *Internals) Modify() {
	i.Modified = time.Now().UTC().Unix()
}