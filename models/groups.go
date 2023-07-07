package models

import "github.com/google/uuid"

type Group struct {
	ID          string   `json:"id"`
	Admin       string   `json:"admin"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Region      string   `json:"region"`
	Members     []string `json:"members"`
}

func NewGroup(admin, region, name, descr string) *Group {
	return &Group{
		ID:          uuid.NewString(),
		Admin:       admin,
		Name:        name,
		Description: descr,
		Region:      region,
	}
}
