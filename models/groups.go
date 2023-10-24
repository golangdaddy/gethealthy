package models

import "github.com/google/uuid"

type Group struct {
	Meta        Internals
	Options     GroupOptions
	ID          string    `json:"id" firestore:"id"`
	Region      string    `json:"region" firestore:"region"`
	Admins      []UserRef `json:"admins" firestore:"admins"`
	Moderators  []UserRef `json:"moderators" firestore:"moderators"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	Email       string    `json:"email" firestore:"email"`
	Website     string    `json:"website" firestore:"website"`
	Socials
}

type GroupOptions struct {
	// controls whether anyone can instantly join a group
	Open bool
	// has threads
	Threads bool
	// has meetings
	Meetings bool
	// has events
	Events bool
}

func NewGroup(user *User, region, name, descr string, options GroupOptions) *Group {
	return &Group{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Region:      region,
		Admins:      []UserRef{user.Ref()},
		Name:        name,
		Description: descr,
		Options:     options,
	}
}

type Meeting struct {
	Meta    Internals
	Options GroupOptions
	ID      string  `json:"id" firestore:"id"`
	Title   string  `json:"title" firestore:"title"`
	Address string  `json:"address" firestore:"address"`
	Cost    float64 `json:"cost" firestore:"cost"`
	Start   string  `json:"start" firestore:"start"`
	// number of minutes
	Duration int `json:"duration" firestore:"duration"`
}

func (group *Group) NewMeeting(title, address string, cost float64, start string, duration int) *Meeting {
	meeting := &Meeting{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Title:    title,
		Address:  address,
		Cost:     cost,
		Start:    start,
		Duration: duration,
	}
	meeting.Meta.Parent = group.ID
	return meeting
}
