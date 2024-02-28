package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
	// Connect to the database
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get the environment variables
	db_uri := os.Getenv("MONGODB_URI")
	if db_uri == "" {
		log.Fatal("MONGODB_URI is nil")
	}

	// Set client options
	clientOption := options.Client().ApplyURI(db_uri)

	// Connect to the MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
