package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	MetaMaskID string `bson:"meta_id" json:"meta_id"`
	Perm       int    `bson:"permission" json:"permission"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	MetaMaskID string             `bson:"meta_id" json:"meta_id"`
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
