package models

import "github.com/google/uuid"

type Group struct {
	ID          string `json:"id" firestore:"id"`
	Region      string `json:"region" firestore:"region"`
	Admin       string `json:"admin" firestore:"admin"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	Email       string `json:"email" firestore:"email"`
	Website     string `json:"website" firestore:"website"`
	// socials
	Members []string `json:"members" firestore:"members"`
}

func NewGroup(admin, region, name, descr string) *Group {
	return &Group{
		ID:          uuid.NewString(),
		Region:      region,
		Admin:       admin,
		Name:        name,
		Description: descr,
	}
}
