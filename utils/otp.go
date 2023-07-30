package utils

import (
	"context"
	"net/http"

	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/cloudfunc"
	"github.com/sirchorg/go/common"
)

const (
	CONST_COL_SESSION = "sessions"
	CONST_COL_OTP     = "otp"
	CONST_COL_USER    = "users"
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
	doc, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	otpRecord := &models.OTP{}
	if err := doc.DataTo(&otpRecord); err != nil {
		return nil, err
	}

	// delete the OTP record
	if _, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Delete(ctx); err != nil {
		return nil, err
	}

	return otpRecord, nil
}

// GetOTP gets OTP record from firestore
func DebugGetOTP(app *common.App, r *http.Request) (*models.OTP, error) {

	ctx := context.Background()

	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}
	id := app.SeedDigest(otp)

	// fetch the OTP record
	doc, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	otpRecord := &models.OTP{}
	if err := doc.DataTo(&otpRecord); err != nil {
		return nil, err
	}

	return otpRecord, nil
}

func CreateSessionSecret(app *common.App, otp *models.OTP) (string, int64, error) {

	ctx := context.Background()

	secret := app.Token256()
	hashedSecret := app.SeedDigest(secret)

	session := models.NewSession(otp.Username)

	// create the firestore session record
	if _, err := app.Firestore().Collection(CONST_COL_SESSION).Doc(hashedSecret).Set(ctx, session); err != nil {
		return "", 0, err
	}

	return secret, session.Expires, nil
}

func GetSessionUser(app *common.App, r *http.Request) (*models.User, error) {

	ctx := context.Background()

	apiKey := r.Header.Get("Authorization")
	id := app.SeedDigest(apiKey)

	// fetch the Session record
	doc, err := app.Firestore().Collection(CONST_COL_SESSION).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	session := &models.Session{}
	if err := doc.DataTo(&session); err != nil {
		return nil, err
	}

	// fetch the user record
	doc, err = app.Firestore().Collection(CONST_COL_USER).Doc(session.Username).Get(ctx)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return user, nil
}
