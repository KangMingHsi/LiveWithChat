package authentication

import "github.com/dgrijalva/jwt-go"

// UserClaims is for jwt to create token.
type UserClaims struct {
	jwt.StandardClaims
	ID  MemberID
}

// TokenService is to generate and verify token.
type TokenService interface {
	// Generate creates token.
	Generate(user *User) (string, error)

	// Verify checks token.
	Verify(accessToken string) (*UserClaims, error)
}
