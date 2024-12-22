package models

type Post struct {
	Message string `bson:"message" json:"message"`
}
