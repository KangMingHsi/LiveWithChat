package jwt

import (
	"errors"
	"stream_subsystem"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ErrUnexpectedMethod is returned wrong claim created method.
var ErrUnexpectedMethod = errors.New("unexpected token signing method")

// ErrInvalidClaims is returned wrong claims.
var ErrInvalidClaims = errors.New("invalid token claims")

// userClaims is for jwt to create token.
type userClaims struct {
	jwt.StandardClaims
	Email  string
	RoleLevel   string
}

func (c *userClaims) GetKey() interface{} {
	return c.Audience
}

func (c *userClaims) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"UserID": c.Audience,
		"Email": c.Email,
		"RoleLevel": c.RoleLevel,
		"IssuedAt": c.IssuedAt,
	}
}

type tokenManager struct {
	secretKey     string
}

// Verify checks token.
func (manager *tokenManager) Verify(accessToken string) (stream_subsystem.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedMethod
			}

			claim := token.Claims.(*userClaims)
			if time.Now().Unix() > claim.ExpiresAt {
				return nil, stream_subsystem.ErrExpiredClaims
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

// NewTokenManager creates a instance of JWTManager.
func NewTokenManager(
		secretKey string) stream_subsystem.TokenManager {
	return &tokenManager{
		secretKey: secretKey,
	}
}
