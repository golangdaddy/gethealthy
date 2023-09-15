package models

import (
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type UserRef struct {
	ID       string
	Username string
}

func DemoUser() *User {
	return NewUser("john@doe.com", "john doe")
}

func NewUser(email, username string) *User {
	user := &User{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Email:    email,
		Username: username,
	}
	user.Profiles.Business.Meta = user.Meta
	user.Profiles.Practitioner.Meta = user.Meta
	user.Profiles.Personal.Meta = user.Meta
	return user
}

type User struct {
	Meta Internals
	ID   string
	// user (0) or practitioner (1) or business (2)
	Account  int      `json:"account" firestore:"account"`
	Email    string   `json:"email" firestore:"email"`
	Username string   `json:"username" firestore:"username"`
	Profiles Profiles `json:"profiles" firestore:"profiles"`
}

func (user *User) Ref() UserRef {
	return UserRef{
		ID:       user.ID,
		Username: user.Username,
	}
}

func (user *User) IsValid() bool {
	log.Println(user.Username)

	if len(user.Username) < 6 {
		return false
	}
	if len(user.Username) > 24 {
		return false
	}
	if strings.Contains(user.Username, " ") {
		return false
	}
	if !isAlphanumeric(strings.Replace(user.Username, "_", "", -1)) {
		return false
	}
	return true
}

func isAlphanumeric(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}
