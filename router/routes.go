package router

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(client *mongo.Client, ctx context.Context) *echo.Echo {
	e := echo.New()
	database := client.Database("chronologger")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/register", controllers.SignUp())
	e.GET("/login", controllers.SignIn())
	e.GET("/logout", controllers.Logout(database))

	e.POST("/register", controllers.Register(database, ctx))
	e.POST("/login", controllers.Login(database, ctx))

	e.GET("/profile", controllers.Profile(database), jwtAuthenticationMiddleware)

	e.Static("/static", "assets")

	e.GET("/", controllers.Home(database))

	templates := e.Group("/templates")
	templates.GET("", controllers.Templates(database))
	templates.GET("/:id", controllers.Template(database))
	templates.GET("/:id/modal", controllers.Modal(database))
	templates.GET("/:id/start", controllers.Start(database))
	templates.POST("/:id/response", controllers.Response(database, client))
	templates.POST("/dismiss", controllers.DismissModal())

	return e
}
