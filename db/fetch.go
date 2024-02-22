package db

import (
	"context"

	"github.com/hay-i/chronologger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDefaultTemplates(ctx context.Context, database *mongo.Database) []models.Template {
	collection := database.Collection("templates")

	var results []models.Template

	filter := bson.M{"default": true}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results
}

func GetTemplate(ctx context.Context, database *mongo.Database, id string) models.Template {
	collection := database.Collection("templates")

	var result models.Template
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	if err = collection.FindOne(ctx, filter).Decode(&result); err != nil {
		panic(err)
	}

	return result
}
