package models

import (
	"github.com/google/uuid"
	"github.com/sirchorg/go/common"
)

type Mail struct {
	Meta       Internals
	ID         string    `json:"id" firestore:"id"`
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Subject    string    `json:"subject" firestore:"subject"`
	Body       string    `json:"body" firestore:"body"`
}

func (user *User) NewMail(app *common.App, subject, body string, recipients ...*User) *Mail {
	return &Mail{
		Meta:       NewInternals(),
		ID:         uuid.NewString(),
		Sender:     user.Ref(),
		Recipients: Users(recipients).Refs(),
		Subject:    subject,
		Body:       body,
	}
}
