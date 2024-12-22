package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPosts struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"userid" json:"author_id,omitempty"`
	Content string             `bson:"message" json:"content,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
