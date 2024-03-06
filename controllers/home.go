package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"
	"github.com/labstack/echo/v4"
)

func Home(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		var isLoggedIn bool
		cookie, err := c.Cookie("token")
		if err != nil {
			isLoggedIn = false
		} else {
			_, err = parseToken(cookie.Value)

			if err != nil {
				isLoggedIn = false
			} else {
				isLoggedIn = true
			}
		}

		component := components.Home(isLoggedIn)

		return render(c, component)
	}
}
