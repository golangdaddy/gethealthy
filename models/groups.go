package models

import "github.com/google/uuid"

type Group struct {
	Meta        Internals
	ID          string       `json:"id" firestore:"id"`
	Options     GroupOptions `json:"options" firestore:"options"`
	Region      string       `json:"region" firestore:"region"`
	Admins      []UserRef    `json:"admins" firestore:"admins"`
	Moderators  []UserRef    `json:"moderators" firestore:"moderators"`
	Name        string       `json:"name" firestore:"name"`
	Description string       `json:"description" firestore:"description"`
	Email       string       `json:"email" firestore:"email"`
	Website     string       `json:"website" firestore:"website"`
	Socials
}

type GroupOptions struct {
	// controls whether anyone can instantly join a group
	Open bool `json:"open" firestore:"open"`
	// has threads
	Threads bool `json:"threads" firestore:"threads"`
	// has meetings
	Meetings bool `json:"meetings" firestore:"meetings"`
	// has events
	Events bool `json:"events" firestore:"events"`
}

func NewPrivateGroup(user *User, region, name, descr string) *Group {
	return &Group{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Region:      region,
		Admins:      []UserRef{user.Ref()},
		Name:        name,
		Description: descr,
	}
}

func NewPublicGroup(user *User, region, name, descr string) *Group {
	group := &Group{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Region:      region,
		Admins:      []UserRef{user.Ref()},
		Name:        name,
		Description: descr,
	}
	group.Options.Open = true
	return group
}
