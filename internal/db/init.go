package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize(ctx context.Context) (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL environment variable not set")
	}

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(err)
	}

	database := client.Database("gooal")

	createCollections(ctx, database)

	return client, err
}
