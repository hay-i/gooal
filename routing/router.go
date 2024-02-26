package routing

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/db"
	"github.com/hay-i/chronologger/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(e *echo.Echo, client *mongo.Client) {
	database := client.Database("chronologger")

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
		success := c.QueryParam("success")
		component := components.Template(template, answers, success == "true")

		return component.Render(requestContext, c.Response().Writer)
	})

	e.GET("templates/:id/start", func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		component := components.Start(template)

		return component.Render(requestContext, c.Response().Writer)
	})

	e.POST("templates/:id/response", func(c echo.Context) error {
		answersCollection := database.Collection("answers")
		templateId := c.Param("id")
		req := c.Request()
		requestContext := req.Context()

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		session, err := client.StartSession()
		if err != nil {
			panic(err)
		}
		defer session.EndSession(requestContext)
		err = session.StartTransaction()
		if err != nil {
			panic(err)
		}
		defer session.AbortTransaction(requestContext)
		defer session.CommitTransaction(requestContext)

		templateObjectId, err := primitive.ObjectIDFromHex(templateId)
		if err != nil {
			panic(err)
		}

		for questionId, value := range req.Form {
			questionObjectId, err := primitive.ObjectIDFromHex(questionId)
			if err != nil {
				panic(err)
			}
			_, insertErr := answersCollection.InsertOne(requestContext, models.Answer{
				TemplateID: templateObjectId,
				QuestionID: questionObjectId,
				Answer:     value[0],
			})

			if insertErr != nil {
				panic(insertErr)
			}
		}

		// Ideally I wanted to send in http.StatusCreated, but it seems that the redirect doesn't work with that status code
		// See: https://github.com/labstack/echo/issues/229#issuecomment-1518502318
		return c.Redirect(http.StatusFound, "/templates/"+templateId+"?success=true")
	})
}
