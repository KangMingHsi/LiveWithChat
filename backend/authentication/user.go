package authentication

import (
	"errors"
	"fmt"
	"net"
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
	IsOnline bool
	IsBlocked bool
	IpAddr  net.IPAddr
	LastLoginTime time.Time
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
		IsOnline: false,
		IsBlocked: false,
		IpAddr: net.IPAddr{},
		LastLoginTime: time.Now(),
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
		IsOnline: user.IsOnline,
		IsBlocked: 		user.IsBlocked,
		IpAddr: user.IpAddr,
		LastLoginTime: user.LastLoginTime,
	}
}

// Login changes user to online state.
func (user *User) Login(ipAddr string) error {
	user.IsOnline = true
	user.IpAddr = net.IPAddr{
		IP: net.ParseIP(ipAddr),
	}
	user.LastLoginTime = time.Now()
	return nil
}

// Logout changes user to offline state.
func (user *User) Logout() error {
	user.IsOnline = false
	return nil
}

// Change user password.
func (user *User) ChangePassword (newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("cannot hash password: %w", err)
    }

	user.HashedPassword = string(hashedPassword)
	return user.Logout()
}

// UserRepository provides access a user store.
type UserRepository interface {
	Store(user *User) error
	Find(id MemberID) (*User, error)
	FindAll() []*User
}

// ErrUnknownUser is used when a user could not be found.
var ErrUnknownUser = errors.New("unknown user")
// ErrWrongPassward is returned when the password is not matched.
var ErrWrongPassward = errors.New("wrong password")
// ErrUserIsBlocked is returned when user is blocked by server.
var ErrUserIsBlocked = errors.New("user is blocked")

// NextMemberID generates a new member ID.
// TODO: Move to infrastructure(?)
func NextMemberID() MemberID {
	return MemberID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}