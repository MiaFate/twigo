package bd

import (
	"context"
	"fmt"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(id string, page int64, search string, typeUser string) ([]*models.Usuario, bool) {
	//ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var result []*models.Usuario
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSkip((page - 1) * 20)
	filter := bson.M{
		"nombre": bson.M{"$regex": "(?i)" + search},
	}
	// filter := bson.D{
	// 	{Key: "nombre", Value: bson.M{"$regex": "(?i)" + search}},
	// }

	//cursor, err := col.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{Key: "name", Value: 1}}))
	cursor, err := col.Find(context.TODO(), filter, opts)
	fmt.Println(err)
	if err != nil {
		return result, false
	}

	var include bool
	for cursor.Next(context.TODO()) {
		var s models.Usuario

		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println("decode = " + err.Error())
			return result, false
		}

		var r models.Relationship
		r.UserId = id
		r.FriendId = s.Id.Hex()

		include = false

		found := GetRelationship(r)
		fmt.Println(found)
		if typeUser == "new" && !found {
			include = true
		}
		if typeUser == "follow" && found {
			include = true
		}
		if r.FriendId == id {
			include = false
		}
		if include {
			s.Password = ""
			result = append(result, &s)
		}
	}
	err = cursor.Err()
	if err != nil {
		return result, false
	}

	cursor.Close(context.TODO())
	//fmt.Println(result)
	return result, true

}
