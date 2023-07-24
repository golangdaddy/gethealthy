package models

import "time"

type Session struct {
	Username string
	Expires  int64
}

func NewSession(username string) *Session {
	return &Session{
		Username: username,
		Expires:  time.Now().UTC().Unix(),
	}
}
