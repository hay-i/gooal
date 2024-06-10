package router

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/controllers"
	"github.com/hay-i/gooal/views"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(client *mongo.Client) *echo.Echo {
	e := echo.New()

	database := client.Database("gooal")

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

	templates := e.Group("/templates", controllers.JwtAuthenticationMiddleware)
	templates.GET("/build", controllers.Build())
	templates.GET("/get-input", controllers.Input())
	templates.DELETE("/delete-input", controllers.DeleteInput())
	templates.POST("/save", controllers.Save(database, client))

	templates.GET("/:id/complete", controllers.CompleteTemplate(database))
	templates.POST("/:id/complete", controllers.Complete(database))

	return e
}
