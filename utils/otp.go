package utils

import (
	"context"
	"net/http"

	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/cloudfunc"
	"github.com/sirchorg/go/common"
)

// GetOTPUser does get OTP record from firestore
func GetOTP(app *common.App, r *http.Request) (*models.OTP, error) {
	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}
	firestoreDoc := app.Firestore().Collection("otp").Doc(otp)
	doc, err := firestoreDoc.Get(context.Background())
	if err != nil {
		return nil, err
	}
	otpRecord := &models.OTP{}
	return otpRecord, doc.DataTo(&otpRecord)
}
