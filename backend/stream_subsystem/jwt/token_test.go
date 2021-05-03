package jwt

import (
	"stream_subsystem"
	"testing"
)

const (
	secret = "secret"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIyNEZEMEVENS04QUIyLTRCMjgtOUQ4OS02ODQwMjc5MTc4QUEiLCJleHAiOjE2MTg5Mzk5MzQsImlhdCI6MTYxODkzOTYzNCwic3ViIjoibm9ybWFsIiwiRW1haWwiOiJ0ZXN0QGxpdmV3aXRoY2hhdC5jb20iLCJSb2xlIjoibm9ybWFsIn0.nNmB9JwBD1j25LO7qODjYFumGrwJKZsJI05B_piiuso"
)

var tm = NewTokenManager(secret)

func TestVerifyExpired(t *testing.T) {
	_, err := tm.Verify(token)

	if err.Error() != stream_subsystem.ErrExpiredClaims.Error() {
		t.Errorf("the token should be expired")
	}
}