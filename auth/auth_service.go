package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/middleware"
)

type AuthService interface {
	GetAccessToken(userID uint, role string) (string, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) GetAccessToken(userID uint, role string) (string, error) {
	claims := &middleware.JwtCustomClaims{
		ID:   userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
