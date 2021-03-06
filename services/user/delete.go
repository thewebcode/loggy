package user

import (
	"context"

	"github.com/jz222/loggy/libs/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func Delete(filter bson.M) (int64, error) {
	collection := mongodb.GetClient().Collection(mongodb.Users)

	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
