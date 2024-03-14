package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/views"
	"github.com/labstack/echo/v4"
)

func JwtAuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			views.AddFlash(c, "You must be logged in to access that page")

			return redirect(c, "/login")
		}

		tokenString := cookie.Value
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.SecretKey), nil
		})

		if err != nil || !token.Valid {
			views.AddFlash(c, "Invalid or expired token")

			return redirect(c, "/login")
		}

		return next(c)
	}
}