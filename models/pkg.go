package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirchorg/go/cloudfunc"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

func AssertKeyValue(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, ok := m[key].(string)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

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
	ID        string   `json:"id" firestore:"id"`
	Email     string   `json:"email" firestore:"email"`
	Name      string   `json:"name" firestore:"name"`
	Avatar    string   `json:"avatar" firestore:"avatar"`
	Following []string `json:"following" firestore:"following"`
	OTP       string   `json:"-" firestore:"otp"`
	Timestamp int64    `json:"timestamp" firestore:"timestamp"`
}
