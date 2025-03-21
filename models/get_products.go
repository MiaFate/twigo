package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProducts struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Product string             `bson:"product" json:"product,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
