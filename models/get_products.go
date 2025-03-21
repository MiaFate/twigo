package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProducts struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Content string             `bson:"message" json:"content,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
