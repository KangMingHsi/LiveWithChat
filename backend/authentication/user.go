package authentication

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

// MemberID uniquely identifies a particular user.
type MemberID string

// User is the central class in domain model
type User struct {
	ID  		MemberID
	HashedPassword    string
	ExpiredTime time.Time
}

// NewUser creates a new user.
func NewUser(id MemberID, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("cannot hash password: %w", err)
    }
	user := &User{
		ID: id,
		HashedPassword: string(hashedPassword),
		ExpiredTime: time.Now(),
	}
	return user, nil
}

// IsCorrectPassword checks password is matched or not.
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
    return err == nil
}

// Clone copies the same instance of user.
func (user *User) Clone() *User {
	return &User{
		ID: 				user.ID,
		HashedPassword:     user.HashedPassword,
		ExpiredTime: 		user.ExpiredTime,
	}
}

// UserRepository provides access a user store.
type UserRepository interface {
	Store(user *User) error
	Find(id MemberID) (*User, error)
	FindAll() []*User
}

// ErrUnknownUser is used when a user could not be found.
var ErrUnknownUser = errors.New("unknown user")

// NextMemberID generates a new member ID.
// TODO: Move to infrastructure(?)
func NextMemberID() MemberID {
	return MemberID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}