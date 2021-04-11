package jwt

import (
	"auth_subsystem"
	"testing"
	"time"
)

const (
	secret = "secret"
	email = "test@livewithchat.com"
	role = "customer"
)

var tm = NewTokenManager(secret, time.Second * 5)

func TestGenerate(t *testing.T) {
	id := auth_subsystem.NextMemberID()

	_, err := tm.Generate(id, email, role)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVerifyPassed(t *testing.T) {
	id := auth_subsystem.NextMemberID()
	accessToken, err := tm.Generate(id, email, role)
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
