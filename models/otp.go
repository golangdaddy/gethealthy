package models

import (
	"context"
	"time"

	"github.com/sirchorg/go/common"
)

type OTP struct {
	Email     string `json:"email" firestore:"email"`
	Username  string `json:"username" firestore:"username"`
	Timestamp int64  `json:"timestamp" firestore:"timestamp"`
}

func NewOTP(email, username string) *OTP {
	return &OTP{
		Email:     email,
		Username:  username,
		Timestamp: time.Now().UTC().Unix(),
	}
}

func (otp *OTP) User(app *common.App) (*ForumUser, error) {
	doc, err := app.Firestore().Collection("users").Doc(otp.Username).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &ForumUser{}
	return user, doc.DataTo(user)
}
