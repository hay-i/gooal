package main

import (
	"time"

	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/hay-i/chronologger/db"
)

func getAddress(c echo.Context) error {
	fake := faker.New()
	title := fake.Address().Address()
	component := address(title)

	return component.Render(c.Request().Context(), c.Response().Writer)
}

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
	// result isn't used anymore, but this can be used as an example for the moment

	component := page()
	e := echo.New()

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/address", getAddress)

	e.Logger.Fatal(e.Start(":1323"))
}
