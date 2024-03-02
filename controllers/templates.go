package controllers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/db"
	"github.com/hay-i/chronologger/models"
	"github.com/hay-i/chronologger/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Home(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		templates := db.GetDefaultTemplates(requestContext, database)
		component := components.Home(templates)

		return component.Render(requestContext, c.Response().Writer)
	}
}

func Template(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		answers := db.GetAnswers(requestContext, database, id)
		success := c.QueryParam("success")
		component := components.Template(template, answers, success == "true")

		return component.Render(requestContext, c.Response().Writer)
	}
}

func Start(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		id := c.Param("id")
		template := db.GetTemplate(requestContext, database, id)
		component := components.Start(template)

		return component.Render(requestContext, c.Response().Writer)
	}
}

func Respond(client *mongo.Client, database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
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
	}
}

func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SignInForm()

		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}

// SignIn will be executed after SignInForm submission.
func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Load our "test" user.
		storedUser := user.LoadTestUser()
		// Initiate a new User struct.
		u := new(user.User)
		// Parse the submitted data and fill the User struct with the data from the SignIn form.
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Compare the stored hashed password, with the hashed version of the password that was received.
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			// If the two passwords don't match, return a 401 status.
			return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
		}
		// If password is correct, generate tokens and set cookies.
		err := auth.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}

		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	}
}
