package main

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/db"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, dbErr := db.Initialize(ctx)
	database := client.Database("chronologger")
	defer func() {
		if dbErr = client.Disconnect(ctx); dbErr != nil {
			panic(dbErr)
		}
	}()
	defer cancel()

	db.Seed(ctx, database)

	e := echo.New()
	e.Static("/static", "assets")
	e.GET("/", func(c echo.Context) error {
		component := components.Home()

		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/template", func(c echo.Context) error {
		templates := db.GetDefaultTemplates(ctx, database)
		component := components.BuildTemplates(templates)

		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
