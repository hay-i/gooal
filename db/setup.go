package db

import (
	"context"
	"time"

	model "github.com/hay-i/chronologger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize() (context.Context, *mongo.Client, error, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOpts := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
        // TODO: Something better than panic
		panic(err)
	}

	return ctx, client, err, cancel
}

func Seed(ctx context.Context, client *mongo.Client) {
    collection := client.Database("chronologger").Collection("templates")

    defaultTemplate := model.DefaultTemplate{ID: "finance", Title: "Default Template #1", CreatedAt: time.Now()}
    defaultTemplateTwo := model.DefaultTemplate{ID: "social", Title: "Default Template #2", CreatedAt: time.Now()}

	collection.InsertOne(ctx, defaultTemplate)
	collection.InsertOne(ctx, defaultTemplateTwo)
}

func GetDefaultTemplates(ctx context.Context, client *mongo.Client) []model.Template {
    collection := client.Database("chronologger").Collection("templates")

    var results []model.Template
    // findOptions := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})
    // collection.FindOne(ctx, bson.D{}, findOptions).Decode(&results)
    cursor, err := collection.Find(ctx, bson.D{})

    if err != nil {
        // TODO: Something better than panic
        panic(err)
    }

    if err = cursor.All(ctx, &results); err != nil {
        // TODO: Something better than panic
        panic(err)
    }

    return results
}
