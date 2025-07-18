package repositories

import (
	"authenticator/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// inserting a new user into the database
func CreateUser(user *models.User) (*models.User, error) {
	// inorderto get the 'users' collection from my GetCollection helper function
	collection := GetCollection("users")

	// creating a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// inserting the user document into the collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// finding a single user by their email address
func GetUserByEmail(email string) (*models.User, error) {
	collection := GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {

		return nil, err // when no document is found, this error should come
	}

	return &user, nil
}
