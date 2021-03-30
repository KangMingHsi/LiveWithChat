package authentication

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	var (
		id = NextMemberID()
		password = "1234"
	)

	user, err := NewUser(id, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IsCorrectPassword(password) {
		t.Errorf("%s is not correct password", password)
	}

	if user.IsBlocked {
		t.Errorf("user shouldn't be blocked")
	}
}

func TestChangePassword(t * testing.T) {
	var (
		password = "1234"
		user = User{}
	)

	err := user.ChangePassword(password)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IsCorrectPassword(password) {
		t.Errorf("%s is not correct password", password)
	}
}

func TestLogin(t *testing.T) {
	var (
		user = User{}
		ipString = "127.0.0.1"
	)

	err := user.Login(ipString)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IsOnline {
		t.Errorf("Login failed")
	}

	if user.IpAddr != ipString {
		t.Errorf("Ip address is not correct")
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

	user, err := NewUser(id, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	cloneUser := user.Clone()
	if (cloneUser.HashedPassword != user.HashedPassword ||
			cloneUser.ID != user.ID ||
			cloneUser.IsOnline != user.IsOnline ||
			cloneUser.IsBlocked != user.IsBlocked ||
			cloneUser.LastLoginTime != user.LastLoginTime ||
			cloneUser.IpAddr != user.IpAddr) {
		t.Errorf("%s", "clone should create identical instance")
	}
}

func TestConvertToMap(t *testing.T) {
	var (
		user *User
		id = NextMemberID()
		password = "1234"
	)

	user, err := NewUser(id, password)
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