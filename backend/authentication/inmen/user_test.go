package inmen

import (
	"authentication"
	"testing"
)

func TestStore(t *testing.T) {
	var (
		user1 = &authentication.User{Email: "a@a.com"}
		user2 = &authentication.User{Email: "b@b.com"}
	)

	r := NewUserRepository()
	err := r.Store(user1)
	if err != nil {
		t.Error(err)
	}

	err = r.Store(user2)
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	var (
		user1 = &authentication.User{Email: "a@a.com"}
		user2 = &authentication.User{Email: "b@b.com"}
	)

	r := NewUserRepository()
	_ = r.Store(user1)
	_ = r.Store(user2)

	dbUser1, err := r.Find("a@a.com")
	if err != nil {
		t.Error(err)
	}

	if dbUser1.Email != "a@a.com" {
		t.Errorf("Email should be the same")
	}

	_, err = r.Find("a@b.com")
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}

func TestFindAll(t *testing.T) {
	var (
		user1 = &authentication.User{Email: "a@a.com"}
		user2 = &authentication.User{Email: "b@b.com"}
	)

	r := NewUserRepository()
	_ = r.Store(user1)
	_ = r.Store(user2)

	dbUsers := r.FindAll()
	if len(dbUsers) != 2 {
		t.Errorf("There should be two users")
	}
}