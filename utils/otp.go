package utils

import (
	"context"
	"net/http"

	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/cloudfunc"
	"github.com/sirchorg/go/common"
)

// GetOTP gets OTP record from firestore
func GetOTP(app *common.App, r *http.Request) (*models.OTP, error) {

	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}

	hashedOTP := app.SeedDigest(otp)

	doc, err := app.Firestore().Collection("otp").Doc(hashedOTP).Get(context.Background())
	if err != nil {
		return nil, err
	}

	otpRecord := &models.OTP{}
	return otpRecord, doc.DataTo(&otpRecord)
}
