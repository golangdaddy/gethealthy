package models

import "github.com/google/uuid"

type Group struct {
	Meta Internals
	ID   string `json:"id" firestore:"id"`
	// controls whether anyone can instantly join a group
	Open        bool     `json:"open" firestore:"open"`
	Region      string   `json:"region" firestore:"region"`
	Admin       string   `json:"admin" firestore:"admin"`
	Moderators  []string `json:"moderators" firestore:"moderators"`
	Name        string   `json:"name" firestore:"name"`
	Description string   `json:"description" firestore:"description"`
	Email       string   `json:"email" firestore:"email"`
	Website     string   `json:"website" firestore:"website"`
}

func NewPrivateGroup(admin, region, name, descr string) *Group {
	return &Group{
		ID:          uuid.NewString(),
		Region:      region,
		Admin:       admin,
		Name:        name,
		Description: descr,
	}
}

func NewPublicGroup(admin, region, name, descr string) *Group {
	return &Group{
		ID:          uuid.NewString(),
		Open:        true,
		Region:      region,
		Admin:       admin,
		Name:        name,
		Description: descr,
	}
}
