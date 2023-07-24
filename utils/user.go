package utils

import (
	"context"
	"fmt"

	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/common"
	"google.golang.org/api/iterator"
)

func GetUserByEmail(app *common.App, email string) (*models.User, error) {

	iter := app.Firestore().Collection("users").Where("email", "==", email).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		user := &models.User{}
		return user, doc.DataTo(user)
	}

	return nil, fmt.Errorf("no user forund via email: %s", email)
}
