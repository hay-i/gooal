package db

import (
	"context"

	"github.com/hay-i/chronologger/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveTemplate(database *mongo.Database, ctx context.Context, formValues models.Template) {
	templateCollection := database.Collection("templates")

	for _, defaultTemplate := range []models.Template{} {
		filter := bson.M{"title": defaultTemplate.Title}
		count, err := templateCollection.CountDocuments(ctx, filter)
		if err != nil {
			panic(err)
		}

		if count > 0 {
			continue
		}

		_, err = templateCollection.InsertOne(ctx, defaultTemplate)
		if err != nil {
			panic(err)
		}
	}
}
