package bd

import (
	"github.com/miafate/twigo/models"
	"golang.org/x/net/context"
)

func DeleteRelationship(t models.Relationship) (bool, error) {
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relationship")

	_, err := col.DeleteOne(context.TODO(), t)
	if err != nil {
		return false, err
	}
	return true, nil
}
