package jwt

import (
	"authentication"
	"errors"
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
	Role   string
}

func (c *userClaims) GetKey() interface{} {
	return c.Email
}

func (c *userClaims) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"UserID": authentication.MemberID(c.Audience),
		"Email": c.Email,
		"Role": c.Role,
		"IssuedAt": c.IssuedAt,
	}
}

type tokenManager struct {
	secretKey     string
	accessTokenDuration time.Duration
	refreshTokenDuration time.Duration
}

func (manager *tokenManager) Generate(
		id authentication.MemberID, email, role string) (
		accessTokenString, newRefreshTokenString string, err error) {
	accessTokenString, err = createToken(
		id, email, role,
		manager.accessTokenDuration,
		false, manager.secretKey)
	if err != nil {
		return "", "", err
	}

	newRefreshTokenString, err = createToken(
		id, email, role,
		manager.refreshTokenDuration,
		true, manager.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, newRefreshTokenString, nil
}

// Verify checks token.
func (manager *tokenManager) Verify(accessToken string, isRefresh bool) (authentication.UserClaims, error) {
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
				return nil, authentication.ErrExpiredClaims
			}

			if !isRefresh && claim.Subject == "refresh" {
				return nil, ErrInvalidClaims
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
				 email, role string,
				 tokenDuration time.Duration,
				 isRefresh bool,
				 secretKey string) (string, error) {
	
	var subject string
	if isRefresh {
		subject = "refresh"
	} else {
		subject = "normal"
	}

	claim := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			IssuedAt: time.Now().Unix(),
			Subject: subject,
			Audience: string(id),
		},
		Email: email,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
