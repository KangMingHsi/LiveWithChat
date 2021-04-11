package auth_subsystem

import (
	"testing"
)

const (
	email = "test@livewithchat.com"
	password = "1234"
	gender = "male"
	nickname = "bot"
	role = "normal"
)

var id = NextMemberID()

func TestConstruction(t *testing.T) {
	_, err := NewUser(id, role, email, gender, nickname, password)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestLogin(t *testing.T) {
	var (
		ipString = "127.0.0.1"
	)
	user, err := NewUser(id, role, email, gender, nickname, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = user.Login(password, ipString)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IsOnline {
		t.Errorf("Login failed")
	}

	if len(user.IpAddr) == 0 {
		t.Errorf("Ip address should be added")
	}
}

func TestChangePassword(t * testing.T) {
	var (
		user = User{}
		ipString = "127.0.0.1"
	)

	err := user.ChangePassword(password)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = user.Login(password, ipString)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestLogout(t *testing.T) {
	var (
		user = User{}
	)

	err := user.Logout()
	if err != nil {
		t.Errorf("%v", err)
	}

	if user.IsOnline {
		t.Errorf("Logout failed")
	}
}

func TestClone(t *testing.T) {
	var (
		user *User
		id = NextMemberID()
		password = "1234"
	)

	user, err := NewUser(id, role, email, gender, nickname, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	cloneUser := user.Clone()
	if (cloneUser.HashedPassword != user.HashedPassword ||
			cloneUser.ID != user.ID ||
			cloneUser.IsOnline != user.IsOnline ||
			cloneUser.IsBlocked != user.IsBlocked ||
			cloneUser.LoginTime != user.LoginTime) {
		t.Errorf("%s", "clone should create identical instance")
	}
}

func TestConvertToMap(t *testing.T) {
	var (
		user *User
		id = NextMemberID()
		password = "1234"
	)

	user, err := NewUser(id, role, email, gender, nickname, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	userMap := user.ConvertToMap()
	if _, ok := userMap["ID"]; !ok {
		t.Errorf("Failed to convert. No ID")
	}

	if v, ok := userMap["ID"].(MemberID); !ok || v != id {
		t.Errorf("Failed to convert. Wrong type or value")
	}
}