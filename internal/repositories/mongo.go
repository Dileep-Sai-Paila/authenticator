package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

// establishh a connection to the MongoDB database.
func ConnectDB(connectionString string) {
	var err error

	// set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// creating a context with a timeout to handle connection attempts.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to MongoDB
	DBClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// just to ping the primary inorderto verify that the connection is alive or not
	err = DBClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {

	collection := DBClient.Database("authenticatorDB").Collection(collectionName) // the database where the 'users' collection will be stored
	return collection                                                             //returns a handle to a specific collection in the database
}
