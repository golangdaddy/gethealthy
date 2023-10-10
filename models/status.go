package models

import "github.com/google/uuid"

type Content struct {
	Type   string `json:"type" firestore:"type"`
	String string `json:"string" firestore:"string"`
}

type Status struct {
	Meta    Internals
	ID      string    `json:"id" firestore:"id"`
	Content []Content `json:"content" firestore:"content"`
}

func (user *User) NewStatus(content ...Content) *Status {
	return &Status{
		Meta:    NewInternals(),
		ID:      uuid.NewString(),
		Content: content,
	}
}
