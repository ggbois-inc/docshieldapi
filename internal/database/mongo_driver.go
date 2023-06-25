package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection
var documents *mongo.Collection
var ctx context.Context = context.TODO()

func init() {
	godotenv.Load()
	uri := os.Getenv("MONGO")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Loaded")
	users = client.Database("blockdb").Collection("users")
	documents = client.Database("blockdb").Collection("documents")
}
