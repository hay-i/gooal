package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/labstack/echo/v4"
)

func JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		currentURL := c.Request().URL.RequestURI()

		if err != nil {
			flash.Add(c, "You must be logged in to access that page", flash.Error)

			pathToReturnTo := fmt.Sprintf("/login?return_to=%s", currentURL)

			// TODO: http.StatusSeeOther
			// Old note: I wanted to send in http.StatusCreated, but it seems that the redirect doesn't work with that status code
			// See: https://github.com/labstack/echo/issues/229#issuecomment-1518502318
			return c.Redirect(http.StatusSeeOther, pathToReturnTo)
		}

		tokenString := cookie.Value
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.SecretKey), nil
		})

		if err != nil || !token.Valid {
			flash.Add(c, "Invalid or expired token", flash.Error)

			pathToReturnTo := fmt.Sprintf("/login?return_to=%s", currentURL)

			return c.Redirect(http.StatusSeeOther, pathToReturnTo)
		}

		return next(c)
	}
}
