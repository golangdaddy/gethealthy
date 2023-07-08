package models

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirchorg/go/cloudfunc"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

type OTP struct {
	Email string
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
	ID        string   `json:"id"`
	Email     string   `json:"_"`
	Name      string   `json:"name"`
	Avatar    string   `json:"avatar"`
	Following []string `json:"following"`
	OTP       string   `json:"_"`
	Timestamp int64    `json:"timestamp"`
}

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
	Type         string         `json:"_"`
	Parent       string         `json:"parent"`
	ID           string         `json:"id"`
	Creator      ForumUser      `json:"creator"`
	Title        string         `json:"title"`
	CountReplies int64          `json:"countReplies"`
	CountViews   int64          `json:"countViews"`
	Replies      []*ThreadReply `json:"replies"`
	Timestamp    int64          `json:"timestamp"`
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
	ID        string    `json:"id"`
	Creator   ForumUser `json:"creator"`
	Content   string    `json:"content"`
	Timestamp int64     `json:"timestamp"`
}

func (reply *ThreadReply) ThreadID() string {
	return strings.Split(reply.ID, "_")[0]
}
