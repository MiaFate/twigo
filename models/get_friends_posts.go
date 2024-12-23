package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFriendsPosts struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId   string             `bson:"userid" json:"userId,omitempty"`
	FriendId string             `bson:"friendid" json:"friendId,omitempty"`
	Post     struct {
		Id      string    `bson:"_id" json:"_id,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		Message string    `bson:"message" json:"message,omitempty"`
	}
}
