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
	e.GET("/login", controllers.SignIn())

	e.POST("/register", controllers.Register(database))
	e.POST("/login", controllers.Login(database))

	e.GET("/logout", controllers.Logout(database), controllers.JwtAuthenticationMiddleware)
	e.GET("/profile", controllers.Profile(database), controllers.JwtAuthenticationMiddleware)
	e.GET("/get-started", controllers.GetStarted(database), controllers.JwtAuthenticationMiddleware)

	questionnaire := e.Group("/questionnaire", controllers.JwtAuthenticationMiddleware)
	questionnaire.POST("/step-one", controllers.StepOne(database))

	e.Static("/static", "assets")

	e.GET("/", controllers.Home(database))

	templates := e.Group("/templates")
	templates.GET("", controllers.Templates(database))
	templates.GET("/:id", controllers.Template(database))
	templates.GET("/:id/modal", controllers.Modal(database))
	templates.GET("/:id/start", controllers.Start(database))
	templates.POST("/:id/response", controllers.Response(database, client))

	return e
}
