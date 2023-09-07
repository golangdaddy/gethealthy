package models

import "github.com/google/uuid"

type Group struct {
	Meta Internals
	ID   string `json:"id" firestore:"id"`
	// controls whether anyone can instantly join a group
	Features struct {
		Open     bool `json:"open" firestore:"open"`
		Forum    bool `json:"forum" firestore:"forum"`
		Meetings bool `json:"meetings" firestore:"meetings"`
		Events   bool `json:"events" firestore:"events"`
	} `json:"features" firestore:"features"`
	Region      string   `json:"region" firestore:"region"`
	Admin       string   `json:"admin" firestore:"admin"`
	Moderators  []string `json:"moderators" firestore:"moderators"`
	Name        string   `json:"name" firestore:"name"`
	Description string   `json:"description" firestore:"description"`
	Email       string   `json:"email" firestore:"email"`
	Website     string   `json:"website" firestore:"website"`
	Socials
}

func NewPrivateGroup(admin, region, name, descr string) *Group {
	return &Group{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Region:      region,
		Admin:       admin,
		Name:        name,
		Description: descr,
	}
}

func NewPublicGroup(admin, region, name, descr string) *Group {
	group := &Group{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Region:      region,
		Admin:       admin,
		Name:        name,
		Description: descr,
	}
	group.Features.Open = true
	return group
}
