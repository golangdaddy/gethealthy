package models

import (
	"github.com/google/uuid"
)

type Mail struct {
	Meta       Internals
	ID         string    `json:"id" firestore:"id"`
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Subject    string    `json:"subject" firestore:"subject"`
	Body       string    `json:"body" firestore:"body"`
}

func (user *User) NewMail(subject, body string, recipients ...*User) *Mail {
	return &Mail{
		Meta:       NewInternals("mail"),
		ID:         uuid.NewString(),
		Sender:     user.Ref(),
		Recipients: Users(recipients).Refs(),
		Subject:    subject,
		Body:       body,
	}
}

type MailReply struct {
	Meta       Internals
	ID         string    `json:"id" firestore:"id"`
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Body       string    `json:"body" firestore:"body"`
}

func (user *User) NewMailReply(op *Mail, body string) *MailReply {
	mail := &MailReply{
		Meta:       NewInternals("mailreply"),
		ID:         uuid.NewString(),
		Sender:     user.Ref(),
		Recipients: op.Recipients,
		Body:       body,
	}
	mail.Meta.Parent = op.ID
	return mail
}
