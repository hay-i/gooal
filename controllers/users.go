package controllers

import (
	"net/http"
	"time"

	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/models"
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
			// TODO: Return a template with the error message
			return c.JSON(http.StatusInternalServerError, "Error while registering")
		}

		expiry, signedToken, err := auth.SignToken(user)
		if err != nil {
			// TODO: Return a template with the error message
			return c.JSON(http.StatusInternalServerError, "Failed to sign token")
		}

		auth.SetCookie(signedToken, expiry, c)

		// TODO: Return a template with the success message
		return c.JSON(http.StatusCreated, "User created")
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
			// TODO: Return a template with the error message
			return c.JSON(http.StatusBadRequest, "Invalid username or password")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			// TODO: Return a template with the error message
			return c.JSON(http.StatusBadRequest, "Invalid username or password")
		}

		expirationTime, signedToken, err := auth.SignToken(user)
		if err != nil {
			// TODO: Return a template with the error message
			return c.JSON(http.StatusInternalServerError, "Failed to sign token")
		}

		auth.SetCookie(signedToken, expirationTime, c)

		// TODO: Return a template with the success message
		return c.JSON(http.StatusOK, echo.Map{"token": signedToken})
	}
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SignUp()

		return render(c, component)
	}
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SignIn()

		return render(c, component)
	}
}

func Profile(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			// TODO: Return a template with the error message
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		tokenString := cookie.Value

		parsedToken, err := auth.ParseToken(tokenString)

		if err != nil {
			// TODO: Return a template with the error message
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		component := components.Profile(parsedToken["sub"].(string))

		return render(c, component)
	}
}

func Logout(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.SetCookie("", time.Now(), c)

		// TODO: Return a template with the success message
		return c.JSON(http.StatusOK, "Logged out")
	}
}
