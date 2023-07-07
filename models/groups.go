package models

import "github.com/google/uuid"

type Group struct {
	ID      string
	Admin   string
	Name    string
	Region  string
	Members []string
}

func NewGroup(admin, name, region string) *Group {
	return &Group{
		ID:     uuid.NewString(),
		Admin:  admin,
		Name:   name,
		Region: region,
	}
}
