package bd

import (
	"context"

	"github.com/miafate/twigo/models"
)

func AddRelationship(t models.Relationship) (bool, error) {
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relationship")

	_, err := col.InsertOne(context.TODO(), t)
	if err != nil {
		return false, err
	}
	return true, nil
}
