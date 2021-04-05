package jwt

import (
	"authentication"
	"testing"
	"time"
)

const (
	secret = "secret"
	email = "test@livewithchat.com"
	role = "customer"
)

var tm = NewTokenManager(secret, time.Second * 5, time.Second * 20)

func TestGenerate(t *testing.T) {
	id := authentication.NextMemberID()

	_, _, err := tm.Generate(id, email, role)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVerifyPassed(t *testing.T) {
	id := authentication.NextMemberID()
	accessToken, refreshToken, err := tm.Generate(id, email, role)
	if err != nil {
		t.Errorf(err.Error())
	}

	claim, err := tm.Verify(accessToken, false)
	if err != nil {
		t.Errorf(err.Error())
	}

	if claim.GetKey() != email {
		t.Errorf("access claim recovers failed")
	}

	claim, err = tm.Verify(refreshToken, true)
	if err != nil {
		t.Errorf(err.Error())
	}

	if claim.GetKey() != email {
		t.Errorf("refresh claim recovers failed")
	}
}
