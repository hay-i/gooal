package main

import (
	"github.com/labstack/echo/v4"

	db "github.com/hay-i/chronologger/db"
	"github.com/hay-i/chronologger/models"
)

func getTemplates(c echo.Context, templates []models.Template) error {
	component := buildTemplates(templates)

	return component.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	ctx, client, dbErr, cancel := db.Initialize()
	db.Seed(ctx, client)

	templates := db.GetDefaultTemplates(ctx, client)

	component := page()
	e := echo.New()

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/template", func(c echo.Context) error {
		return getTemplates(c, templates)
	})

	e.Logger.Fatal(e.Start(":1323"))

	// TODO: do we need a teardown function?
	defer func() {
		if dbErr = client.Disconnect(ctx); dbErr != nil {
			// TODO: Something better than panic
			panic(dbErr)
		}
	}()

	defer cancel()
}
