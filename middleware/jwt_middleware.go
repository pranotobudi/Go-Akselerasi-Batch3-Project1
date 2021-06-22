package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type JwtCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(key),
	})
}
