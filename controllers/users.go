package controllers

import (
	"time"

	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/models"
	"github.com/hay-i/chronologger/views"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		collection := database.Collection("users")
		var user models.User
		if err := c.Bind(&user); err != nil {
			return err
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		_, err := collection.InsertOne(requestContext, user)
		if err != nil {
			views.AddFlash(c, "Error while registering", views.FlashError)

			return redirect(c, "/register")
		}

		expiry, signedToken, err := auth.SignToken(user)
		if err != nil {
			views.AddFlash(c, "Error while registering", views.FlashError)

			return redirect(c, "/register")
		}

		auth.SetCookie(signedToken, expiry, c)
		views.AddFlash(c, "You have successfully registered", views.FlashSuccess)

		return redirect(c, "/")
	}
}

func Login(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		collection := database.Collection("users")
		var credentials models.User
		if err := c.Bind(&credentials); err != nil {
			return err
		}

		var user models.User
		err := collection.FindOne(requestContext, bson.M{"username": credentials.Username}).Decode(&user)
		if err != nil {
			views.AddFlash(c, "Invalid username or password", views.FlashError)

			return redirect(c, "/login")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			views.AddFlash(c, "Invalid username or password", views.FlashError)

			return redirect(c, "/login")
		}

		expirationTime, signedToken, err := auth.SignToken(user)
		if err != nil {
			views.AddFlash(c, "Error while logging in", views.FlashError)

			return redirect(c, "/login")
		}

		auth.SetCookie(signedToken, expirationTime, c)
		views.AddFlash(c, "You have successfully logged in", views.FlashSuccess)

		return redirect(c, "/")
	}
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SignUp()

		return renderWithoutNav(c, component)
	}
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SignIn()

		return renderWithoutNav(c, component)
	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.SetCookie("", time.Now(), c)

		views.AddFlash(c, "You have successfully logged out", views.FlashSuccess)

		return redirect(c, "/")
	}
}

func Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		component := components.Profile(parsedToken["sub"].(string))

		return renderBase(c, component)
	}
}
