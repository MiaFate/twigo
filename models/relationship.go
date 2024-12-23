package models

type Relationship struct {
	UserId             string `bson:"userid" json:"userId"`
	UserRelationshipId string `bson:"userfriendid" json:"userFriendId"`
}
