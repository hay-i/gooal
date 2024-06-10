package db

import (
	"context"
	"time"

	"github.com/hay-i/gooal/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// Make a generic one and move this to model
func SaveTemplate(database *mongo.Database, template models.Template) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	templateCollection := database.Collection("templates")

	_, err := templateCollection.InsertOne(ctx, template)
	if err != nil {
		panic(err)
	}
}
