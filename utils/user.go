package utils

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/golangdaddy/gethealthy/models"
	"github.com/sirchorg/go/common"
	"google.golang.org/api/iterator"
)

// UserCollection abstracts the handling of subdata to within the user objectz
func UserCollection(app *common.App, user *models.User, collectionID string) *firestore.CollectionRef {
	return app.Firestore().Collection("users").Doc(user.ID).Collection(collectionID)
}

func GetUser(app *common.App, ref models.UserRef) (*models.User, error) {
	doc, err := app.Firestore().Collection("users").Doc(ref.ID).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	return user, doc.DataTo(user)
}

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
