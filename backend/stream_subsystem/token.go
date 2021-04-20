package stream_subsystem

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
	Generate(id, email, role string) (string, error)

	// Verify checks token.
	Verify(token string) (UserClaims, error)

	// Verify checks without expire time.
	VerifyWithoutExpired(token string) (UserClaims, error)
}

// ErrExpiredClaims is returned expired claims.
var ErrExpiredClaims = errors.New("token is expired")
// ErrNoIssuedAt is returned when there is no issued_at in claim.
var ErrNoIssuedAt = errors.New("no issued at")
