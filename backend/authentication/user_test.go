package authentication

import "testing"

func TestConstruction(t *testing.T) {
	id := NextMemberID()
	password := "1234"
	user, err := NewUser(id, password)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !user.IsCorrectPassword(password) {
		t.Errorf("%s is not correct password", password)
	}
}
