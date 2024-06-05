package router

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/controllers"
	"github.com/hay-i/chronologger/views"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(client *mongo.Client, ctx context.Context) *echo.Echo {
	e := echo.New()

	database := client.Database("chronologger")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(views.SessionStore))

	e.GET("/register", controllers.SignUp())
	e.POST("/register", controllers.Register(database))
	e.GET("/login", controllers.SignIn())
	e.POST("/login", controllers.Login(database))

	e.GET("/logout", controllers.Logout(), controllers.JwtAuthenticationMiddleware)
	e.GET("/profile", controllers.Profile(), controllers.JwtAuthenticationMiddleware)

	questionnaire := e.Group("/questionnaire", controllers.JwtAuthenticationMiddleware)
	questionnaire.GET("/step-one", controllers.StepOne())
	questionnaire.GET("/step-two", controllers.StepTwo())

	e.Static("/static", "assets")

	e.GET("/", controllers.Home())

	templates := e.Group("/templates")
	templates.GET("/build", controllers.Build())
	templates.GET("/builder", controllers.Builder())
	templates.POST("/save", controllers.Save(database, client, ctx))

	return e
}
