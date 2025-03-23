package models

type Product struct {
	Product string `bson:"product" json:"product,omitempty"`
}
