package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Get this from an environment variable
// generate some sort of csrf token
// OpenSSL package to generate
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

		expirationTime := time.Now().Add(1 * time.Hour)
		claims := &jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(SecretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to sign token")
		}

		return c.JSON(http.StatusOK, echo.Map{"token": signedToken})
	}
}
