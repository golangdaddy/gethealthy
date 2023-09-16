package models

import "strconv"

func (user *User) GetUsernameRef() *Username {
	ref := &Username{
		ID:    user.ID,
		Index: map[string][]string{},
	}
	max := len(user.Username)
	if max > 14 {
		max = 14
	}
	for x := 3; x < max; x++ {
		ref.Index[strconv.Itoa(x)] = []string{user.Username[:x]}
	}
	return ref
}

type Username struct {
	ID    string
	Index map[string][]string
}