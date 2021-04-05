package authentication

import "errors"

// UserClaims is to store data of token
type UserClaims interface {
	// Get key
	GetKey () interface{}

	// Convert to map
	ConvertToMap () map[string]interface{}
}

// TokenManager is to generate and verify token.
type TokenManager interface {
	// Generate creates token.
	Generate(id MemberID, email, role string) (string, string, error)

	// Verify checks token.
	Verify(accessToken string, isRefresh bool) (UserClaims, error)
}

// ErrExpiredClaims is returned expired claims.
var ErrExpiredClaims = errors.New("token is expired")