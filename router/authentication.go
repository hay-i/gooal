package router

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/controllers"
	"github.com/labstack/echo/v4"
)

func jwtAuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		tokenString := cookie.Value
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(controllers.SecretKey), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		return next(c)
	}
}
