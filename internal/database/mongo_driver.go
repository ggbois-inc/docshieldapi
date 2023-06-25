package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection
var documents *mongo.Collection
var ctx context.Context = context.TODO()

type Permission struct {
	MetaMaskID string `bson:"meta_id" json:"meta_id"`
	Perm       int    `bson:"permission" json:"permission"`
}

type Document struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Filename    string             `bson:"filename" json:"filename"`
	CID         string             `bson:"cid" json:"cid"`
	ShortCode   string             `bson:"short" json:"short"`
	CreatedBy   string             `bson:"meta_id" json:"meta_id"`
	CreatedOn   time.Time          `bson:"created_on" json:"created_on"`
	Permissions []Permission       `bson:"permissions" json:"permissions"`
}

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

func GetDocumentByCode(shortcode string) Document {
	var result Document
	err := documents.FindOne(ctx, bson.D{primitive.E{Key: "short", Value: shortcode}}).Decode(&result)
	if err == nil {
		return result
	}
	return Document{}
}
