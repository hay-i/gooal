package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/hay-i/gooal/internal/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPOST(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		returnToParam := c.QueryParam("return_to")
		returnTo := "/"
		var returnToQueryParam string
		if returnToParam != "" {
			returnTo = returnToParam
			returnToQueryParam = "?return_to=" + returnTo
		}

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
			flash.Add(c, "Error while registering", flash.Error)

			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/register%s", returnToQueryParam))
		}

		expiry, signedToken, err := auth.SignToken(user)
		if err != nil {
			flash.Add(c, "Error while registering", flash.Error)

			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/register%s", returnToQueryParam))
		}

		auth.SetCookie(signedToken, expiry, c)
		flash.Add(c, "You have successfully registered", flash.Success)

		return c.Redirect(http.StatusSeeOther, returnTo)
	}
}

func LoginPOST(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		returnToParam := c.QueryParam("return_to")
		returnTo := "/"
		var returnToQueryParam string
		if returnToParam != "" {
			returnTo = returnToParam
			returnToQueryParam = "?return_to=" + returnTo
		}

		requestContext := c.Request().Context()
		collection := database.Collection("users")
		var credentials models.User
		if err := c.Bind(&credentials); err != nil {
			return err
		}

		var user models.User
		err := collection.FindOne(requestContext, bson.M{"username": credentials.Username}).Decode(&user)
		if err != nil {
			flash.Add(c, "Invalid username or password", flash.Error)

			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/login%s", returnToQueryParam))
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			flash.Add(c, "Invalid username or password", flash.Error)

			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/login%s", returnToQueryParam))
		}

		expirationTime, signedToken, err := auth.SignToken(user)
		if err != nil {
			flash.Add(c, "Error while logging in", flash.Error)

			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/login%s", returnToQueryParam))
		}

		auth.SetCookie(signedToken, expirationTime, c)
		flash.Add(c, "You have successfully logged in", flash.Success)

		return c.Redirect(http.StatusSeeOther, returnTo)
	}
}

func SignUpGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		returnToParam := c.QueryParam("return_to")
		var returnTo string
		if returnToParam != "" {
			returnTo = "?return_to=" + returnToParam
		}

		component := components.SignUp(returnTo)

		return controllers.RenderWithoutNav(c, component)
	}
}

func SignInGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		returnToParam := c.QueryParam("return_to")
		var returnTo string
		if returnToParam != "" {
			returnTo = "?return_to=" + returnToParam
		}

		component := components.SignIn(returnTo)

		return controllers.RenderWithoutNav(c, component)
	}
}

func LogoutGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.SetCookie("", time.Now(), c)

		flash.Add(c, "You have successfully logged out", flash.Success)

		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func ProfileGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := auth.GetTokenFromCookie(c)
		if err != nil {
			return auth.HandleInvalidToken(c, "You must be logged in to access that page")
		}

		parsedToken, err := auth.ParseToken(tokenString)
		if err != nil {
			return auth.HandleInvalidToken(c, "Invalid or expired token")
		}

		username := auth.TokenToUsername(parsedToken)
		component := components.Profile(username)

		return controllers.RenderBase(c, component)
	}
}
