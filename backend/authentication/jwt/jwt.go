package jwt

import (
	"authentication"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ErrUnexpectedMethod is returned wrong claim created method.
var ErrUnexpectedMethod = errors.New("unexpected token signing method")

// ErrInvalidClaims is returned wrong claims.
var ErrInvalidClaims = errors.New("invalid token claims")

// JWTService is to generate and verify token.
type JWTService struct {
	secretKey     string
	tokenDuration time.Duration
}

func (manager *JWTService) Generate(user *authentication.User) (string, error) {
	claims := authentication.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		ID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify checks token.
func (manager *JWTService) Verify(accessToken string) (*authentication.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&authentication.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedMethod
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid token: %s", err))
	}

	claims, ok := token.Claims.(*authentication.UserClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

// NewJWTService creates a instance of JWTManager.
func NewJWTService(secretKey string, tokenDuration time.Duration) authentication.TokenService {
	return &JWTService{
		secretKey,
		tokenDuration,
	}
}
