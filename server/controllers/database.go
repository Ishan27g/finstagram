package controllers

import (
"context"
"log"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"

)

// Client is exported
var Client *mongo.Client

// Connect to the MongoDB database
func ConnectDatabase() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	Client = client
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected.")
}
