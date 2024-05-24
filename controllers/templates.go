package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/db"
	"github.com/hay-i/chronologger/models"

	"github.com/hay-i/chronologger/views"
	"github.com/labstack/echo/v4"
)

func MyTemplates(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		cookie, err := c.Cookie("token")
		if err != nil {
			views.AddFlash(c, "You must be logged in to access that page", views.FlashError)

			return redirect(c, "/login")
		}

		tokenString := cookie.Value

		parsedToken, err := auth.ParseToken(tokenString)

		if err != nil {
			views.AddFlash(c, "Invalid or expired token", views.FlashError)

			return redirect(c, "/login")
		}

		username := parsedToken["sub"].(string)

		templates := db.GetMyTemplates(requestContext, database, username)

		component := components.Templates(templates)

		return renderBase(c, component)
	}
}

func Templates(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		templates := db.GetDefaultTemplates(requestContext, database)
		component := components.Templates(templates)

		return renderBase(c, component)
	}
}

func Template(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		answers := db.GetAnswers(requestContext, database, id)

		component := components.Template(template, answers)

		return renderBase(c, component)
	}
}

func Modal(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		component := components.Modal(template)

		return renderNoBase(c, component)
	}
}

func Start(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		component := components.Start(template)

		return renderBase(c, component)
	}
}

func Response(database *mongo.Database, client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		answersCollection := database.Collection("answers")
		templateId := c.Param("id")
		req := c.Request()
		requestContext := req.Context()

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		dbSession, err := client.StartSession()
		if err != nil {
			panic(err)
		}
		defer dbSession.EndSession(requestContext)
		err = dbSession.StartTransaction()
		if err != nil {
			panic(err)
		}
		defer dbSession.AbortTransaction(requestContext)
		defer dbSession.CommitTransaction(requestContext)

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

		views.AddFlash(c, "Your response has been saved", views.FlashSuccess)

		return redirect(c, "/templates/"+templateId)
	}
}

func Build(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		goal := c.QueryParam("goal")
		focus := c.QueryParam("focus")

		component := components.Build(goal, focus)

		return renderNoBase(c, component)
	}
}
