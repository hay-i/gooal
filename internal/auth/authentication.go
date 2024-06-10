package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/labstack/echo/v4"
)

func IsLoggedIn(c echo.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		return false
	} else {
		_, err = ParseToken(cookie.Value)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func HandleInvalidToken(c echo.Context, message string) error {
	flash.Add(c, message, flash.Error)

	return c.Redirect(http.StatusSeeOther, "/login")
}

func TokenToUsername(token jwt.MapClaims) string {
	return token["sub"].(string)
}
