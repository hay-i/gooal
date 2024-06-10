package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/gooal/auth"
	"github.com/hay-i/gooal/views"
	"github.com/labstack/echo/v4"
)

func JwtAuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			views.AddFlash(c, "You must be logged in to access that page", views.FlashError)

			// TODO: http.StatusSeeOther
			// Old note: I wanted to send in http.StatusCreated, but it seems that the redirect doesn't work with that status code
			// See: https://github.com/labstack/echo/issues/229#issuecomment-1518502318
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		tokenString := cookie.Value
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.SecretKey), nil
		})

		if err != nil || !token.Valid {
			views.AddFlash(c, "Invalid or expired token", views.FlashError)

			return c.Redirect(http.StatusSeeOther, "/login")
		}

		return next(c)
	}
}
