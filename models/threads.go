package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (user *ForumUser) NewThread(parent, title string) *Thread {
	return &Thread{
		ID:        uuid.NewString(),
		Creator:   *user,
		Title:     title,
		Parent:    parent,
		Timestamp: getTime(),
	}
}

type Thread struct {
	Type         string         `json:"-" firestore:"-"`
	Parent       string         `json:"parent" firestore:"parent"`
	ID           string         `json:"id" firestore:"id"`
	Creator      ForumUser      `json:"creator" firestore:"creator"`
	Title        string         `json:"title" firestore:"title"`
	CountReplies int64          `json:"countReplies" firestore:"countReplies"`
	CountViews   int64          `json:"countViews" firestore:"countviews"`
	Replies      []*ThreadReply `json:"replies" firestore:"replies"`
	Timestamp    int64          `json:"timestamp" firestore:"timestamp"`
}

func (thread *Thread) Reply(user *ForumUser, content string) *ThreadReply {
	return &ThreadReply{
		ID:        fmt.Sprintf("%s_%s", thread.ID, uuid.NewString()),
		Creator:   *user,
		Content:   content,
		Timestamp: getTime(),
	}
}

type ThreadReply struct {
	ID        string    `json:"id" firestore:"id"`
	Creator   ForumUser `json:"creator"  firestore:"creator"`
	Content   string    `json:"content" firestore:"content"`
	Timestamp int64     `json:"timestamp" firestore:"timestamp"`
}

func (reply *ThreadReply) ThreadID() string {
	return strings.Split(reply.ID, "_")[0]
}
