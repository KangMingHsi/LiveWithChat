package auth

import (
	"authentication"
	"errors"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrWrongPassward is returned when the password is not matched.
var ErrWrongPassward = errors.New("wrong password")

// Service is the interface that provides authorization methods.
type Service interface {
	// Login checks user email and password, returns token.
	Login(id authentication.MemberID, password string) (string, error)

	// Logout.
	Logout(id authentication.MemberID) error

	// Register creates new user.
	Register(password string) (authentication.MemberID, error)
}

type service struct {
	users authentication.UserRepository
	jwtService authentication.TokenService
}

func (s *service) Login(id authentication.MemberID, password string) (string, error) {
	if id == "" || password == "" {
		return "", ErrInvalidArgument
	}

	user, err := s.users.Find(id)
	if err != nil {
		return "", err
	}

	if !user.IsCorrectPassword(password) {
		return "", ErrWrongPassward
	}

	return s.jwtService.Generate(user)
}

func (s *service) Logout(id authentication.MemberID) error {
	return nil
}

func (s *service) Register(password string) (authentication.MemberID, error) {
	if password == "" {
		return "", ErrInvalidArgument
	}
	id := authentication.NextMemberID()
	user, err := authentication.NewUser(id, password)
	if err != nil {
		return "", err
	}

	if err := s.users.Store(user); err != nil {
		return "", err
	}

	return user.ID, nil
}

// NewService creates a auth service with necessary dependencies.
func NewService(users authentication.UserRepository, jwtService authentication.TokenService) Service {
	return &service{
		users: users,
		jwtService: jwtService,
	}
}