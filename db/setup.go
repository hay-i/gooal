package db

import (
	"context"
	"time"

	"github.com/hay-i/chronologger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize(ctx context.Context) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(err)
	}

	database := client.Database("chronologger")

	collections, err := database.ListCollectionNames(ctx, bson.M{"name": "templates"})
	if err != nil {
		panic(err)
	}

	if len(collections) > 0 {
		return client, nil
	}

	// TODO: Add Default boolean, and default it to false
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

	return client, err
}

func Seed(ctx context.Context, database *mongo.Database) {
	collection := database.Collection("templates")

	defaultTemplates := []models.Template{
		{Title: "Default Template #1", Description: "My description 1", CreatedAt: time.Now(), Default: true},
		{Title: "Default Template #2", Description: "My description 2", CreatedAt: time.Now(), Default: true},
		{Title: "Default Template #3", Description: "My description 3", CreatedAt: time.Now(), Default: true},
	}

	for _, defaultTemplate := range defaultTemplates {
		filter := bson.M{"title": defaultTemplate.Title}
		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			panic(err)
		}

		if count > 0 {
			continue
		}

		_, err = collection.InsertOne(ctx, defaultTemplate)
		if err != nil {
			panic(err)
		}
	}
}

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

// TODO: Extract to another file
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
