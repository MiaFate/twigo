package models

type Image struct {
	Avatar string `bson:"avatar" json:"avatar,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
}
