package models

type Relationship struct {
	UserId   string `bson:"userid" json:"userId"`
	FriendId string `bson:"friendid" json:"friendId"`
}
