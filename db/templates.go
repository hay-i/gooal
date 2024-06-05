package db

import (
	"context"

	"github.com/hay-i/gooal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func SaveTemplate(database *mongo.Database, ctx context.Context, template models.Template) {
	templateCollection := database.Collection("templates")

	_, err := templateCollection.InsertOne(ctx, template)
	if err != nil {
		panic(err)
	}
}
