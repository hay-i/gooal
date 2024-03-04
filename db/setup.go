package db

import (
	"context"

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

	createCollections(ctx, database)

	return client, err
}

func Seed(ctx context.Context, database *mongo.Database) {
	seedTemplates(ctx, database)
	seedAnswers(ctx, database)
}
