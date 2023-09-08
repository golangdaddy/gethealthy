package models

import "github.com/google/uuid"

type Topic struct {
	Meta        Internals
	ID          string
	Name        string
	Description string
}

func NewTopic(name, description string) *Topic {
	return &Topic{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}
