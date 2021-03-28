package authentication

// UserClaims is to store data of token
type UserClaims interface {
	// Get user id
	GetID () MemberID
}

// TokenManager is to generate and verify token.
type TokenManager interface {
	// Generate creates token.
	Generate(id MemberID) (string, string, error)

	// Verify checks token.
	Verify(accessToken string) (UserClaims, error)

	// Refresh accessToken.
	Refresh(refreshToken string) (string, string, error)
}
