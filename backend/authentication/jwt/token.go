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
	tokenDuration time.Duration
}

func (manager *tokenManager) Generate(
		id authentication.MemberID, email, role string) (
		accessTokenString string, err error) {
	accessTokenString, err = createToken(
		id, email, role,
		manager.tokenDuration,
		manager.secretKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
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

			claim := token.Claims.(*userClaims) 
			if time.Now().Unix() > claim.ExpiresAt {
				return nil, authentication.ErrExpiredClaims
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

// Verify checks without expire time.
func (manager *tokenManager) VerifyWithoutExpired(tokenString string) (authentication.UserClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &userClaims{})
	if err != nil {
        return nil, err
    }

	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, ErrUnexpectedMethod
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

// NewTokenManager creates a instance of JWTManager.
func NewTokenManager(
		secretKey string,
		tokenDuration time.Duration) authentication.TokenManager {
	return &tokenManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

func createToken(id authentication.MemberID,
				 email, role string,
				 tokenDuration time.Duration,
				 secretKey string) (string, error) {
	claim := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			IssuedAt: time.Now().Unix(),
			Subject: "normal",
			Audience: string(id),
		},
		Email: email,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
