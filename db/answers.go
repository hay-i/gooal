package db

import (
	"context"
	"time"

	"github.com/hay-i/gooal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func SaveAnswer(database *mongo.Database, answer models.Answer) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	templateCollection := database.Collection("answers")

	_, err := templateCollection.InsertOne(ctx, answer)
	if err != nil {
		panic(err)
	}
}
