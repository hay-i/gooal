package router

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/internal/controllers/answers"
	"github.com/hay-i/gooal/internal/controllers/homes"
	gooal_middleware "github.com/hay-i/gooal/internal/controllers/middleware"
	"github.com/hay-i/gooal/internal/controllers/questionnaires"
	"github.com/hay-i/gooal/internal/controllers/templates"
	"github.com/hay-i/gooal/internal/controllers/users"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(client *mongo.Client) *echo.Echo {
	e := echo.New()

	database := client.Database("gooal")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(flash.SessionStore))

	e.Static("/static", "internal/assets")

	e.GET("/register", users.SignUpGET())
	e.POST("/register", users.RegisterPOST(database))
	e.GET("/login", users.SignInGET())
	e.POST("/login", users.LoginPOST(database))

	e.GET("/logout", users.LogoutGET(), gooal_middleware.JwtAuthentication)
	e.GET("/profile", users.ProfileGET(), gooal_middleware.JwtAuthentication)

	questionnaire := e.Group("/questionnaire", gooal_middleware.JwtAuthentication)
	questionnaire.GET("/step-one", questionnaires.StepOneGET())
	questionnaire.GET("/step-two", questionnaires.StepTwoGET())

	e.GET("/", homes.HomeGET())

	templatesGroup := e.Group("/templates", gooal_middleware.JwtAuthentication)
	templatesGroup.GET("/build", templates.BuildGET())
	templatesGroup.GET("/get-input", templates.InputGET())
	templatesGroup.DELETE("/delete-input", templates.InputDELETE())
	templatesGroup.POST("/save", templates.SavePOST(database, client))

	templatesGroup.GET("/:id/complete", answers.AnswerTemplateGET(database))
	templatesGroup.POST("/:id/complete", answers.AnswerTemplatePOST(database))

	return e
}
