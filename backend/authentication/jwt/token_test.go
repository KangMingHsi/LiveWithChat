package jwt

import (
	"authentication"
	"testing"
	"time"
)

var secret = "secret"
var tm = NewTokenManager(secret, time.Second * 5, time.Second * 20)

func TestGenerate(t *testing.T) {
	id := authentication.NextMemberID()

	_, _, err := tm.Generate(id)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVerifyPassed(t *testing.T) {
	id := authentication.NextMemberID()
	accessToken, refreshToken, err := tm.Generate(id)
	if err != nil {
		t.Errorf(err.Error())
	}

	claim, err := tm.Verify(accessToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if claim.GetID() != id {
		t.Errorf("access claim recovers failed")
	}

	claim, err = tm.Verify(refreshToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if claim.GetID() != id {
		t.Errorf("refresh claim recovers failed")
	}
}

func TestRefresh(t *testing.T) {
	id := authentication.NextMemberID()
	_, refreshToken, err := tm.Generate(id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, _, err = tm.Refresh(refreshToken)
	if err != nil {
		t.Errorf(err.Error())
	}
}