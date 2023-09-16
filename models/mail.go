package models

import (
	"github.com/sirchorg/go/common"
)

type Mail struct {
	Meta       Internals
	ID         string    `json:"id" firestore:"id"`
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Subject    string    `json:"subject" firestore:"subject"`
	Body       string    `json:"body" firestore:"body"`
	Timestamp  int64     `json:"timestamp" firestore:"timestamp"`
}

func (user *User) NewMail(app *common.App, recipient *User) *Mail {
	return &Mail{
		Meta:       NewInternals(),
		ID:         app.Token256(),
		Sender:     user.Ref(),
		Recipients: []UserRef{user.Ref()},
		Timestamp:  app.TimeNow().Unix(),
	}
}
