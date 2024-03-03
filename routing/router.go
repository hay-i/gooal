package routing

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(e *echo.Echo, client *mongo.Client) {
	database := client.Database("chronologger")

	e.Use(middleware.Logger())

	e.Static("/static", "assets")

	e.GET("/", controllers.Home(database))

	e.GET("templates", controllers.Templates(database))
	e.GET("templates/:id", controllers.Template(database))
	e.GET("templates/:id/modal", controllers.Modal(database))
	e.GET("templates/:id/start", controllers.Start(database))
	e.POST("templates/:id/response", controllers.Response(database, client))
	e.POST("templates/dismiss", controllers.DismissModal())
}
