package routing

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(e *echo.Echo, database *mongo.Database) {
	e.Use(middleware.Logger())

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		requestContext := c.Request().Context()
		templates := db.GetDefaultTemplates(requestContext, database)
		component := components.Home(templates)

		return component.Render(requestContext, c.Response().Writer)
	})

	e.GET("templates/:id", func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		answers := db.GetAnswers(requestContext, database, id)
		component := components.Template(template, answers)

		return component.Render(requestContext, c.Response().Writer)
	})

    e.GET("templates/:id/start", func(c echo.Context) error {
        requestContext := c.Request().Context()
        id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		component := components.Start(template)

		return component.Render(requestContext, c.Response().Writer)
    })
}
