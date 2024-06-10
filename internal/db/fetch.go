package db

import (
	"context"
	"time"

	"github.com/hay-i/gooal/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Make this generic and then have a separate function for each type of model
func GetTemplate(database *mongo.Database, id string) models.Template {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
