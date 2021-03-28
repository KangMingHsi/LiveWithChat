package auth

import (
	"errors"
	"testing"

	"authentication"
)

var password = "1234"
var newPassword = "5678"
var repository mockUserRepository
var tokenManager mockTokenManager
var claim mockUserClaim
var s = NewService(&repository, &tokenManager)

func TestRegister(t *testing.T) {
	id, err := s.Register(password)
	claim.id = id
	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != id {
		t.Errorf("user id is not the same %s != %s", user.ID, id)
	}

	if !user.IsCorrectPassword(password) {
		t.Errorf("password is wrong")
	}
}

func TestLogin(t *testing.T) {
	user, err := repository.Find("")
	if err != nil {
		t.Fatal(err)
	}

	accessToken, refreshToken, err := s.Login(user.ID, password, "")
	if err != nil {
		t.Errorf("%v", err)
	}

	if accessToken != "access" {
		t.Errorf("accessToken is wrong")
	}

	if refreshToken != "refresh" {
		t.Errorf("refreshToken is wrong")
	}

	if !user.IsOnline {
		t.Errorf("user is not online")
	}
}

func TestCheckAccessToken(t *testing.T) {
	accessToken, refreshToken, err := s.CheckAndRefresh("access", "refresh")
	if err != nil {
		t.Fatal(err)
	}

	if accessToken != "access" {
		t.Errorf("accessToken is wrong")
	}

	if refreshToken != "refresh" {
		t.Errorf("refreshToken is wrong")
	}
}

func TestRefreshToken(t *testing.T) {
	accessToken, refreshToken, err := s.CheckAndRefresh("invalidAccess", "refresh")
	if err != nil {
		t.Fatal(err)
	}

	if accessToken != "newAccess" {
		t.Errorf("accessToken is wrong")
	}

	if refreshToken != "newRefresh" {
		t.Errorf("refreshToken is wrong")
	}
}

func TestLogout(t *testing.T) {
	err := s.Logout("access")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangePassword(t *testing.T) {
	user, err := repository.Find("")
	if err != nil {
		t.Fatal(err)
	}

	accessToken, _, err := s.Login(user.ID, password, "")
	if err != nil {
		t.Errorf("%v", err)
	}

	err = s.ChangePassword(newPassword, accessToken)
	if err != nil {
		t.Errorf("%v", err)
	}

	if user.IsCorrectPassword(password) {
		t.Errorf("password is %s, not %s", newPassword, password)
	}
}

type mockUserRepository struct {
	user *authentication.User
}

func (r *mockUserRepository) Store(c *authentication.User) error {
	r.user = c
	return nil
}

func (r *mockUserRepository) Find(id authentication.MemberID) (*authentication.User, error) {
	if r.user != nil {
		return r.user, nil
	}
	return nil, authentication.ErrUnknownUser
}

func (r *mockUserRepository) FindAll() []*authentication.User {
	return []*authentication.User{r.user}
}

type mockTokenManager struct {}

func (s *mockTokenManager) Generate(id authentication.MemberID) (string, string, error) {
	return "access", "refresh", nil
}

func (s *mockTokenManager) Verify(accessToken string) (authentication.UserClaims, error) {
	if accessToken != "access" && accessToken != "refresh" {
		return &claim, errors.New("token is invalid")
	}
	return &claim, nil
}

func (s *mockTokenManager) Refresh(refreshToken string) (string, string, error) {
	return "newAccess", "newRefresh", nil
}

type mockUserClaim struct {
	id authentication.MemberID
}

func (c *mockUserClaim) GetID () authentication.MemberID {
	return c.id
}