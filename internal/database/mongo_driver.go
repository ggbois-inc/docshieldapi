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

func GetDocuments(meta_id string) []Document {
	cur, err := documents.Find(ctx, bson.D{primitive.E{Key: "meta_id", Value: meta_id}})
	var docs []Document
	if err = cur.All(context.TODO(), &docs); err != nil {
		panic(err)
	}
	log.Printf("Document searched for %s", meta_id)
	return docs
}

func GetDocumentByCID(cid string) Document {
	var result Document
	err := documents.FindOne(ctx, bson.D{primitive.E{Key: "cid", Value: cid}}).Decode(&result)
	if err == nil {
		return result
	}
	return Document{}
}

func GetDocumentByCode(shortcode string) Document {
	var result Document
	err := documents.FindOne(ctx, bson.D{primitive.E{Key: "short", Value: shortcode}}).Decode(&result)
	if err == nil {
		return result
	}
	return Document{}
}

func CreateDocument(meta_id string, filename string, cid string, shortcode string) Document {
	doc := Document{ID: primitive.NewObjectID(), CreatedBy: meta_id, CID: cid, Filename: filename, ShortCode: shortcode, CreatedOn: time.Now(), Permissions: []Permission{}}
	_, err := documents.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Document created for %s", meta_id)
	return doc
}

func CreateUser(meta_id string) User {
	user := User{ID: primitive.NewObjectID(), MetaMaskID: meta_id}
	var result User
	err := users.FindOne(ctx, bson.D{primitive.E{Key: "meta_id", Value: meta_id}}).Decode(&result)
	if err == nil {
		log.Printf("User %s exists", meta_id)
		return result
	}
	_, err = users.InsertOne(ctx, user)
	log.Printf("User %s created", meta_id)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
