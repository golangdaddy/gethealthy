package models

import "github.com/google/uuid"

type Topic struct {
	Meta        Internals
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func NewTopic(name, description string) *Topic {
	return &Topic{
		Meta:        NewInternals(),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}

type TopicUpdate struct {
	Meta     Internals
	ID       string       `json:"id" firestore:"id"`
	Username string       `json:"username" firestore:"username"`
	Update   string       `json:"update" firestore:"update"`
	Media    []*MediaFile `json:"media" firestore:"media"`
}

func (user *User) NewTopicUpdate(update string, mediaFiles ...*MediaFile) *TopicUpdate {
	return &TopicUpdate{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Username: user.Username,
		Update:   update,
		Media:    mediaFiles,
	}
}

type TopicQuestion struct {
	Meta     Internals
	ID       string `json:"id" firestore:"id"`
	Username string `json:"username" firestore:"username"`
	Question string `json:"question" firestore:"question"`
}

func (user *User) NewTopicQuestion(question string) *TopicQuestion {
	return &TopicQuestion{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Username: user.Username,
		Question: question,
	}
}
