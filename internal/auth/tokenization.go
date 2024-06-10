package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hay-i/gooal/internal/models"
	"github.com/labstack/echo/v4"
)

// TODO: Move Secret to ENV
// https://github.com/hay-i/gooal/issues/35
var SecretKey = "my_secret"

func SignToken(user models.User) (time.Time, string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))

	return expirationTime, signedToken, err
}

func SetCookie(signedToken string, expiry time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = signedToken
	cookie.Expires = expiry
	cookie.HttpOnly = true // Make the cookie inaccessible to JavaScript running in the browser
	// https://github.com/hay-i/gooal/issues/35
	// cookie.Secure = true
	c.SetCookie(cookie)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid or expired token")
	}
}

func GetTokenFromCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func TokenToUsername(token jwt.MapClaims) string {
	return token["sub"].(string)
}
