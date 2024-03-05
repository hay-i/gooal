package routing

import (
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/chronologger/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtAuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
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

func Initialize(e *echo.Echo, client *mongo.Client) {
	database := client.Database("chronologger")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", controllers.Register(database))
	e.POST("/login", controllers.Login(database))

	// e.GET("/profile", Profile, JwtAuthenticationMiddleware)

	e.Static("/static", "assets")

	e.GET("/", controllers.Home(database))

	templates := e.Group("/templates")
	templates.GET("", controllers.Templates(database))
	templates.GET("/:id", controllers.Template(database))
	templates.GET("/:id/modal", controllers.Modal(database))
	templates.GET("/:id/start", controllers.Start(database))
	templates.POST("/:id/response", controllers.Response(database, client))
	templates.POST("/dismiss", controllers.DismissModal())
}
