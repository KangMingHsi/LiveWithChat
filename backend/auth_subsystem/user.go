package auth_subsystem

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// MemberID uniquely identifies a particular user.
type MemberID string

// User is the central class in domain model
type User struct {
	ID  		MemberID

	Email string
	HashedPassword    string
	Gender   string
	Nickname string

	Role string
	IsOnline bool
	IsBlocked bool
	IpAddr  []string

	LimitationPeriod time.Time
	LoginTime time.Time
}

// NewUser creates a new user.
func NewUser(
		id MemberID,
		role, email, gender, nickname, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("cannot hash password: %w", err)
    }
	user := &User{
		ID: id,
		Email: email,
		HashedPassword: string(hashedPassword),
		Gender: gender,
		Nickname: nickname,
		Role: role,
		IsOnline: false,
		IsBlocked: false,
		IpAddr: []string{},
		LimitationPeriod: time.Now(),
		LoginTime: time.Now(),
	}
	return user, nil
}

// Clone copies the same instance of user.
func (user *User) Clone() *User {
	return &User{
		ID: 				user.ID,
		Email: user.Email,
		HashedPassword:     user.HashedPassword,
		Gender: user.Gender,
		Nickname: user.Nickname,
		Role: user.Role,
		IsOnline: user.IsOnline,
		IsBlocked: 		user.IsBlocked,
		IpAddr: user.IpAddr,
		LimitationPeriod: user.LimitationPeriod,
		LoginTime: user.LoginTime,
	}
}

// Login changes user to online state.
func (user *User) Login(password, ipAddr string) error {
	if user.IsBlocked {
		return ErrUserIsBlocked
	}

	if err := user.isCorrectPassword(password); err != nil {
		return err
	}

	user.IsOnline = true
	if idx := sort.SearchStrings(user.IpAddr, ipAddr); idx == len(user.IpAddr) {
		user.IpAddr = append(user.IpAddr, ipAddr)
	}

	user.LoginTime = time.Now()
	return nil
}

// Logout changes user to offline state.
func (user *User) Logout() error {
	user.IsOnline = false
	user.LimitationPeriod = time.Now()
	return nil
}

// Convert to map type
func (user *User) ConvertToMap() map[string]interface{} {
	return map[string]interface{} {
		"ID": user.ID,
		"Email": user.Email,
		"HashedPassword": user.HashedPassword,
		"Gender": user.Gender,
		"Nickname": user.Nickname,
		"Role": user.Role,
		"IsOnline": user.IsOnline,
		"IsBlocked": user.IsBlocked,
		"IpAddr": user.IpAddr,
		"LimitationPeriod": user.LimitationPeriod,
		"LoginTime": user.LoginTime,
	}
}

func (user *User) RoleLevel() int {
	if user.Role == "normal" {
		return 1
	}
	return 0
}

// Change user password.
func (user *User) ChangePassword(newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("cannot hash password: %w", err)
    }

	user.HashedPassword = string(hashedPassword)
	return user.Logout()
}

// isCorrectPassword checks password is matched or not.
func (user *User) isCorrectPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
}

// UserRepository provides access to a user store.
type UserRepository interface {
	Store(user *User) error
	Find(email string) (*User, error)
	FindAll() []*User
}

// ErrUnknownUser is used when a user could not be found.
var ErrUnknownUser = errors.New("unknown user")
// ErrWrongPassward is returned when the password is not matched.
var ErrWrongPassward = errors.New("wrong password")
// ErrUserIsBlocked is returned when user is blocked by server.
var ErrUserIsBlocked = errors.New("user is blocked")
// ErrUserShouldLogin is returned when user is forced to re-login.
var ErrUserShouldLogin = errors.New("user should login again")

// NextMemberID generates a new member ID.
// TODO: Move to infrastructure(?)
func NextMemberID() MemberID {
	return MemberID(strings.ToUpper(uuid.New().String()))
}