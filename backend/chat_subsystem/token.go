package chat_subsystem

import "errors"

// UserClaims is to store data of token
type UserClaims interface {
	// Get key
	GetKey () interface{}

	// Convert to map
	ConvertToMap () map[string]interface{}
}

// TokenManager is to verify token.
type TokenManager interface {
	// Verify checks token.
	Verify(token string) (UserClaims, error)
}

// ErrExpiredClaims is returned expired claims.
var ErrExpiredClaims = errors.New("token is expired")
// ErrNoIssuedAt is returned when there is no issued_at in claim.
var ErrNoIssuedAt = errors.New("no issued at")
