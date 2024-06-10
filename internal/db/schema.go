package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createCollections(ctx context.Context, database *mongo.Database) {
	createTemplateCollection(ctx, database)
	createAnswersCollection(ctx, database)
}

func createAnswersCollection(ctx context.Context, database *mongo.Database) {
	collections, err := database.ListCollectionNames(ctx, bson.M{"name": "answers"})

	if err != nil {
		panic(err)
	}

	if len(collections) > 0 {
		return
	}

	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"template_id", "created_at", "username"},
		"properties": bson.M{
			"template_id": bson.M{
				"bsonType":    "objectId",
				"description": "id of the template the answers belong to, which is required",
			},
			"created_at": bson.M{
				"bsonType":    "date",
				"description": "date the template was created, which is required",
			},
			"username": bson.M{
				"bsonType":    "string",
				"description": "username of the user who answered the template",
			},
		},
	}

	validator := bson.M{"$jsonSchema": jsonSchema}
	opts := options.CreateCollection().SetValidator(validator)

	if err = database.CreateCollection(ctx, "answers", opts); err != nil {
		panic(err)
	}
}

func createTemplateCollection(ctx context.Context, database *mongo.Database) {
	collections, err := database.ListCollectionNames(ctx, bson.M{"name": "templates"})

	if err != nil {
		panic(err)
	}

	if len(collections) > 0 {
		return
	}

	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"title", "description", "created_at", "username"},
		"properties": bson.M{
			"title": bson.M{
				"bsonType":    "string",
				"description": "title of the template, which is required",
			},
			"username": bson.M{
				"bsonType":    "string",
				"description": "username of the user who created the template",
			},
			"description": bson.M{
				"bsonType":    "string",
				"description": "description of the template, which is required",
			},
			"created_at": bson.M{
				"bsonType":    "date",
				"description": "date the template was created, which is required",
			},
		},
	}

	validator := bson.M{"$jsonSchema": jsonSchema}
	opts := options.CreateCollection().SetValidator(validator)

	if err = database.CreateCollection(ctx, "templates", opts); err != nil {
		panic(err)
	}
}
