package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Save(database *mongo.Database, collectionName string, document interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Collection(collectionName)

	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		panic(err)
	}
}
