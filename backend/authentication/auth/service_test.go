package auth

import (
	"testing"

	"authentication"
)

var password = "1234"
var repository mockUserRepository
var tokenService mockTokenService
var s = NewService(&repository, &tokenService)

func TestRegister(t *testing.T) {
	id, err := s.Register(password)
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

	token, err := s.Login(user.ID, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	if token != "pass" {
		t.Errorf("token is wrong")
	}
}

func TestLogout(t *testing.T) {
	err := s.Logout("")
	if err != nil {
		t.Fatal(err)
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

type mockTokenService struct {}

func (s *mockTokenService) Generate(user *authentication.User) (string, error) {
	return "pass", nil
}
func (s *mockTokenService) Verify(accessToken string) (interface{}, error) {
	return nil, nil
}
