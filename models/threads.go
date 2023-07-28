package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (user *User) NewThread(parent, title string) *Thread {
	return &Thread{
		Internals: NewInternals(),
		ID:        uuid.NewString(),
		Username:  user.Username,
		Title:     title,
		Parent:    parent,
		Timestamp: getTime(),
	}
}

type Thread struct {
	Internals
	Type      string `json:"-" firestore:"-"`
	Parent    string `json:"parent" firestore:"parent"`
	ID        string `json:"id" firestore:"id"`
	Username  string `json:"username" firestore:"username"`
	Title     string `json:"title" firestore:"title"`
	Timestamp int64  `json:"timestamp" firestore:"timestamp"`
}

func (thread *Thread) Reply(user *User, content string) *ThreadReply {
	return &ThreadReply{
		ID:        fmt.Sprintf("%s_%s", thread.ID, uuid.NewString()),
		Creator:   *user,
		Content:   content,
		Timestamp: getTime(),
	}
}

type ThreadReply struct {
	ID        string `json:"id" firestore:"id"`
	Creator   User   `json:"creator"  firestore:"creator"`
	Content   string `json:"content" firestore:"content"`
	Timestamp int64  `json:"timestamp" firestore:"timestamp"`
}

func (reply *ThreadReply) ThreadID() string {
	return strings.Split(reply.ID, "_")[0]
}
