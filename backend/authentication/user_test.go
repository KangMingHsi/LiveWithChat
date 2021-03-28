package authentication

import (
	"net"
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
	ipAddr, err := net.ResolveIPAddr("ip", ipString)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IpAddr.IP.Equal(ipAddr.IP) {
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