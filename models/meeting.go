package models

import "github.com/google/uuid"

type Meeting struct {
	Meta    Internals
	Options GroupOptions
	ID      string  `json:"id" firestore:"id"`
	Title   string  `json:"title" firestore:"title"`
	Address string  `json:"address" firestore:"address"`
	Cost    float64 `json:"cost" firestore:"cost"`
	Date    string  `json:"date" firestore:"date"`
	Time    string  `json:"time" firestore:"time"`
	// number of minutes
	Duration int `json:"duration" firestore:"duration"`
}

func (group *Group) NewMeeting(title, address string, cost float64, date, time string, duration int) *Meeting {
	meeting := &Meeting{
		Meta:     NewInternals(),
		ID:       uuid.NewString(),
		Title:    title,
		Address:  address,
		Cost:     cost,
		Date:     date,
		Time:     time,
		Duration: duration,
	}
	meeting.Meta.Parent = group.ID
	return meeting
}

type MeetingComment struct {
	Meta    Internals
	User    UserRef
	ID      string `json:"id" firestore:"id"`
	Content string `json:"content" firestore:"content"`
}

func (meeting *Meeting) NewComment(user *User, content string) *MeetingComment {
	comment := &MeetingComment{
		Meta:    NewInternals(),
		User:    user.Ref(),
		ID:      uuid.NewString(),
		Content: content,
	}
	comment.Meta.Parent = meeting.ID
	return comment
}
