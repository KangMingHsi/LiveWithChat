package jwt

import (
	"auth_subsystem"
	"testing"
	"time"
)

const (
	secret = "secret"
	email = "test@livewithchat.com"
	roleLevel = 1
)

var tm = NewTokenManager(secret, time.Second * 5)

func TestGenerate(t *testing.T) {
	id := auth_subsystem.NextMemberID()

	_, err := tm.Generate(string(id), email, roleLevel)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVerifyPassed(t *testing.T) {
	id := auth_subsystem.NextMemberID()
	accessToken, err := tm.Generate(string(id), email, roleLevel)
	if err != nil {
		t.Errorf(err.Error())
	}

	claim, err := tm.Verify(accessToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if claim.GetKey() != email {
		t.Errorf("access claim recovers failed")
	}
}
