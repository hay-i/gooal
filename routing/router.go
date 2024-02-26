package routing

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(e *echo.Echo, client *mongo.Client) {
	database := client.Database("chronologger")

	e.Use(middleware.Logger())

	e.Static("/static", "assets")
	adminGroup := e.Group("/admin")
	adminGroup.GET("", controllers.Admin())
	// TODO: Deprecated
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecretKey()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	e.GET("/", controllers.Home(database))
	e.GET("templates/:id", controllers.Template(database))
	e.GET("templates/:id/start", controllers.Start(database))
	e.POST("templates/:id/response", controllers.Respond(client, database))
	e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
	e.POST("/user/signin", controllers.SignIn())
}
