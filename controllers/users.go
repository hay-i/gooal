package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// https://github.com/hay-i/chronologger/issues/35
var SecretKey = "my_secret"

// Register handles user registration.
func Register(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		collection := database.Collection("users")
		var user models.User
		if err := c.Bind(&user); err != nil {
			return err
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error while registering")
		}

		expiry, signedToken, err := signToken(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to sign token")
		}

		setCookie(signedToken, expiry, c)

		return c.JSON(http.StatusCreated, "User created")
	}
}

// Login handles user login.
func Login(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		collection := database.Collection("users")
		var credentials models.User
		if err := c.Bind(&credentials); err != nil {
			return err
		}

		var user models.User
		err := collection.FindOne(context.TODO(), bson.M{"username": credentials.Username}).Decode(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid username or password")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid username or password")
		}

		expirationTime, signedToken, err := signToken(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to sign token")
		}

		setCookie(signedToken, expirationTime, c)

		return c.JSON(http.StatusOK, echo.Map{"token": signedToken})
	}
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		component := components.SignUp()

		return component.Render(requestContext, c.Response().Writer)
	}
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestContext := c.Request().Context()
		component := components.SignIn()

		return component.Render(requestContext, c.Response().Writer)
	}
}

// Function to parse, decode, and verify the JWT
func parseToken(tokenString string) jwt.MapClaims {
	// Parse the token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

func Profile(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		tokenString := cookie.Value

		parsedToken := parseToken(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
		}

		requestContext := c.Request().Context()
		component := components.Profile(parsedToken["sub"].(string))

		return component.Render(requestContext, c.Response().Writer)
	}
}

func Logout(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		setCookie("", time.Now(), c)

		return c.JSON(http.StatusOK, "Logged out")
	}
}
