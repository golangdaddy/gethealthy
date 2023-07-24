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

	ctx := context.Background()

	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}
	id := app.SeedDigest(otp)

	// fetch the OTP record
	doc, err := app.Firestore().Collection("otp").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	otpRecord := &models.OTP{}
	if err := doc.DataTo(&otpRecord); err != nil {
		return nil, err
	}

	// delete the OTP record
	if _, err := app.Firestore().Collection("otp").Doc(id).Delete(ctx); err != nil {
		return nil, err
	}

	return otpRecord, nil
}

func CreateSessionSecret(app *common.App, otp *models.OTP) (string, error) {

	ctx := context.Background()

	secret := app.Token256()
	hashedSecret := app.SeedDigest(secret)

	session := models.NewSession(otp.Username)

	// create the firestore session record
	if _, err := app.Firestore().Collection("session").Doc(hashedSecret).Set(ctx, session); err != nil {
		return "", err
	}

	return secret, nil
}
