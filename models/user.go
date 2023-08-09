package models

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func DemoUser() *User {
	return NewUser("john@doe.com", "john doe")
}

// NewInternals returns a boilerplate internal object
func NewInternals() Internals {
	i := Internals{}
	i.Created = time.Now().UTC().Unix()
	return i
}

type Internals struct {
	Moderation struct {
		Blocked  bool
		Approved bool
	}
	Searchable bool
	Created    int64
	Modified   int64
	Stats      struct {
		Followers int64
		Views     int64
		Likes     int64
		Replies   int64
	}
}

func (i *Internals) Modify() {
	i.Modified = time.Now().UTC().Unix()
}

func NewUser(email, username string) *User {
	return &User{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Email:    email,
		Username: username,
	}
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
