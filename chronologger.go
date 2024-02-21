package main

import (
	"time"

	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/hay-i/chronologger/db"
)

func main() {
	ctx, client := db.SetupDb()

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
	e.GET("/", func(c echo.Context) error {
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
