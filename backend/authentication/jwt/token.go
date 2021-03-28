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

// userClaims is for jwt to create token.
type userClaims struct {
	jwt.StandardClaims
	ID  authentication.MemberID
}

func (c *userClaims) GetID () authentication.MemberID {
	return c.ID
}

type tokenManager struct {
	secretKey     string
	accessTokenDuration time.Duration
	refreshTokenDuration time.Duration
}

func (manager *tokenManager) Generate(id authentication.MemberID) (string, string, error) {
	accessTokenString, err := createToken(id, manager.accessTokenDuration, manager.secretKey)
	if err != nil {
		return "", "", err
	}

	newRefreshTokenString, err := createToken(id, manager.refreshTokenDuration, manager.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, newRefreshTokenString, nil
}

// Verify checks token.
func (manager *tokenManager) Verify(accessToken string) (authentication.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedMethod
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		claims, _ := token.Claims.(*userClaims)
		return claims, err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

// Refresh accessToken and refreshToken
func (manager *tokenManager) Refresh(refreshToken string) (string, string, error) {
	claim, err := manager.Verify(refreshToken)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("(RefreshToken) %s", err))
	}

	return manager.Generate(claim.(*userClaims).GetID())
}

// NewTokenManager creates a instance of JWTManager.
func NewTokenManager(secretKey string,
				     accessTokenDuration time.Duration,
				     refreshTokenDuration time.Duration) authentication.TokenManager {
	return &tokenManager{
		secretKey: secretKey,
		accessTokenDuration: accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func createToken(id authentication.MemberID,
				 tokenDuration time.Duration,
				 secretKey string) (string, error) {
	claim := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
		},
		ID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
