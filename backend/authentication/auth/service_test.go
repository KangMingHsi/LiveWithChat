package auth

import (
	"errors"
	"testing"
	"time"

	"authentication"
)

const(
	password = "1234"
	newPassword = "5678"
	email = "test@livewithchat.com"
	nickname = "test"
	gender = "male"
)

var repository mockUserRepository
var cache mockUserRepository
var tokenManager mockTokenManager
var claim mockUserClaim
var s = NewService(&repository, &cache, &tokenManager)

func TestRegister(t *testing.T) {
	id, err := s.Register(
		email, gender, nickname, password)

	claim.email = email
	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.Find(email)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != id {
		t.Errorf("user id is not the same %s != %s", user.ID, id)
	}
}

func TestLogin(t *testing.T) {
	user, err := repository.Find("")
	if err != nil {
		t.Fatal(err)
	}

	accessToken, err := s.Login(user.Email, password, "")
	if err != nil {
		t.Errorf("%v", err)
	}

	if accessToken != "access" {
		t.Errorf("accessToken is wrong")
	}

	if !user.IsOnline {
		t.Errorf("user is not online")
	}
}

func TestCheckAccessToken(t *testing.T) {
	err := s.Check("access")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefreshToken(t *testing.T) {
	accessToken, err := s.Refresh("access")
	if err != nil {
		t.Fatal(err)
	}

	if accessToken != "access" {
		t.Errorf("accessToken is wrong")
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

	accessToken, err := s.Login(user.Email, password, "")
	if err != nil {
		t.Errorf("%v", err)
	}

	err = s.ChangePassword(newPassword, accessToken)
	if err != nil {
		t.Errorf("%v", err)
	}

	if err := user.Login(newPassword, "127.0.0.1"); err != nil {
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

func (r *mockUserRepository) Find(email string) (*authentication.User, error) {
	if r.user != nil {
		return r.user, nil
	}
	return nil, authentication.ErrUnknownUser
}

func (r *mockUserRepository) FindAll() []*authentication.User {
	return []*authentication.User{r.user}
}

type mockTokenManager struct {}

func (s *mockTokenManager) Generate(
		id authentication.MemberID, email, role string) (string, error) {
	return "access", nil
}

func (s *mockTokenManager) Verify(accessToken string,) (authentication.UserClaims, error) {
	if accessToken != "access"{
		return &claim, errors.New("access token is invalid")
	}
	return &claim, nil
}

func (s *mockTokenManager) VerifyWithoutExpired(accessToken string,) (authentication.UserClaims, error) {
	if accessToken != "access"{
		return &claim, errors.New("access token is invalid")
	}
	return &claim, nil
}

type mockUserClaim struct {
	email string
}

func (c *mockUserClaim) GetKey() interface{} {
	return c.email
}

func (c *mockUserClaim) ConvertToMap() map[string]interface{}{
	return map[string]interface{}{
		"UserID": authentication.MemberID("0"),
		"Email": email,
		"Role": "normal",
		"IssuedAt": time.Now().Unix(),
	}
}