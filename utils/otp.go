package utils

import (
	"context"
	"net/http"

	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/cloudfunc"
	"github.com/sirchorg/go/common"
)

func GetOTPUser(app *common.App, r *http.Request) (*models.ForumUser, error) {
	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}

	firestoreDoc := app.Firestore().Collection("otp").Doc(otp)

	// get OTP record from firestore
	doc, err := firestoreDoc.Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &models.ForumUser{}
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	return user, nil
}
