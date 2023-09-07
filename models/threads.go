package models

import (
	"strings"

	"github.com/google/uuid"
)

func (user *User) NewThread(parent, title string) *Thread {
	return &Thread{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Username: user.Username,
		Title:    title,
		Parent:   parent,
	}
}

type Thread struct {
	Meta     Internals
	Type     string `json:"-" firestore:"-"`
	Parent   string `json:"parent" firestore:"parent"`
	ID       string `json:"id" firestore:"id"`
	Username string `json:"username" firestore:"username"`
	Title    string `json:"title" firestore:"title"`
}

func (thread *Thread) Reply(user *User, content string) *ThreadReply {
	return &ThreadReply{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Thread:   thread.ID,
		Title:    thread.Title,
		Username: user.Username,
		Content:  content,
	}
}

type ThreadReply struct {
	Meta     Internals
	ID       string `json:"id" firestore:"id"`
	Thread   string `json:"thread" firestore:"thread"`
	Title    string `json:"title" firestore:"title"`
	Username string `json:"username" firestore:"username"`
	Content  string `json:"content" firestore:"content"`
}

func (reply *ThreadReply) ThreadID() string {
	return strings.Split(reply.ID, "_")[0]
}
