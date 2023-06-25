package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	MetaMaskID string `bson:"meta_id" json:"meta_id"`
	Perm       bool   `bson:"permission" json:"permission"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	MetaMaskID string             `bson:"meta_id" json:"meta_id"`
}

type Document struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Filename         string             `bson:"filename" json:"filename"`
	Filesize         string             `bson:"file_size" json:"file_size"`
	CID              string             `bson:"cid" json:"cid"`
	ShortCode        string             `bson:"short" json:"short"`
	PrivateShortCode string             `bson:"short_private" json:"short_private"`
	CreatedBy        string             `bson:"meta_id" json:"meta_id"`
	CreatedOn        time.Time          `bson:"created_on" json:"created_on"`
	Permissions      []Permission       `bson:"permissions" json:"permissions"`
}
