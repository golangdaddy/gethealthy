package models

import (
	"github.com/sirchorg/go/common"
)

type Mail struct {
	ID        string `json:"id" firestore:"id"`
	Sender    string `json:"sender" firestore:"sender"`
	Recipient string `json:"recipient" firestore:"recipient"`
	Subject   string `json:"subject" firestore:"subject"`
	Body      string `json:"body" firestore:"body"`
}

func (user *User) NewMail(app *common.App, recipient *User) *Mail {
	return &Mail{
		ID:        app.Token256(),
		Sender:    user.Username,
		Recipient: recipient.Username,
	}
}
