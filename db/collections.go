package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createCollections(ctx context.Context, database *mongo.Database) {
	createTemplateCollection(ctx, database)
	createAnswers(ctx, database)
}

func createAnswers(ctx context.Context, database *mongo.Database) {
	collections, err := database.ListCollectionNames(ctx, bson.M{"name": "answers"})

	if err != nil {
		panic(err)
	}

	if len(collections) > 0 {
		return
	}

	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"template_id", "question_id", "answer"},
		"properties": bson.M{
			"template_id": bson.M{
				"bsonType":    "objectId",
				"description": "id of the template the answers belong to, which is required",
			},
			"question_id": bson.M{
				"bsonType":    "objectId",
				"description": "id of the question the answers belong to, which is required",
			},
			"answer": bson.M{
				"bsonType":    "string",
				"description": "answer to the question, which is required",
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
		"required": []string{"title", "description"},
		"properties": bson.M{
			"title": bson.M{
				"bsonType":    "string",
				"description": "title of the template, which is required",
			},
			"description": bson.M{
				"bsonType":    "string",
				"description": "description of the template, which is required",
			},
			"default": bson.M{
				"bsonType":    "bool",
				"description": "whether the template is a default seeded template",
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
