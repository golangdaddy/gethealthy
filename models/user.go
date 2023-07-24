package models

import "github.com/google/uuid"

func DemoUser() *ForumUser {
	return NewForumUser("john@doe.com", "john doe", "https://image.com/a")
}

func NewForumUser(email, name, avatar string) *ForumUser {
	return &ForumUser{
		ID:        uuid.NewString(),
		Email:     email,
		Name:      name,
		Avatar:    avatar,
		Following: []string{},
		Timestamp: getTime(),
	}
}

type ForumUser struct {
	ID string `json:"id" firestore:"id"`
	// user or business
	Type        string   `json:"type" firestore:"type"`
	Email       string   `json:"email" firestore:"email"`
	Name        string   `json:"name" firestore:"name"`
	Firstname   string   `json:"firstname" firestore:"firstname"`
	Lastname    string   `json:"lastname" firestore:"lastname"`
	Description string   `json:"description" firestore:"description"`
	Avatar      string   `json:"avatar" firestore:"avatar"`
	Following   []string `json:"following" firestore:"following"`
	OTP         string   `json:"-" firestore:"otp"`
	Timestamp   int64    `json:"timestamp" firestore:"timestamp"`
}
