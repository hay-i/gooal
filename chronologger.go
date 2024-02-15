package main

import (
	"context"
	"time"

	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(ctx, clientOpts)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	collection := client.Database("chronologger").Collection("goals")
	fake := faker.New()
	title := fake.Address().Address()
	goal := Goal{Title: title, CreatedAt: time.Now()}
	collection.InsertOne(ctx, goal)

	var result Goal
	findOptions := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})
	collection.FindOne(ctx, bson.D{}, findOptions).Decode(&result)

	component := hello(result.Title)
	e := echo.New()

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
