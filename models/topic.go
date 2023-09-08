package models

import "github.com/google/uuid"

type Topic struct {
	Meta        Internals
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func NewTopic(name, description string) *Topic {
	return &Topic{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}
