package auth

import (
	"authentication"
	"errors"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides authorization methods.
type Service interface {
	// Login checks user email and password, returns token.
	Login(id authentication.MemberID, password string, ipAddr string) (string, string, error)

	// Logout.
	Logout(accessToken string) error

	// Register creates new user.
	Register(password string) (authentication.MemberID, error)

	// Check token and refresh token if refreshToken is still valid.
	CheckAndRefresh(accessToken, refreshToken string) (string, string, error)

	// Resets password
	ChangePassword(newPassword, accessToken string) error
}

type service struct {
	users authentication.UserRepository
	tokenManager authentication.TokenManager
}

func (s *service) Login(id authentication.MemberID, password string, ipAddr string) (string, string, error) {
	if id == "" || password == "" {
		return "", "", ErrInvalidArgument
	}

	user, err := s.users.Find(id)
	if err != nil {
		return "", "", err
	}

	if !user.IsCorrectPassword(password) {
		return "", "", authentication.ErrWrongPassward
	}

	if user.IsBlocked {
		return "", "", authentication.ErrUserIsBlocked
	}

	err = user.Login(ipAddr)
	if err != nil {
		return "", "", err
	}

	err = s.users.Store(user)
	if err != nil {
		return "", "", err
	}

	accessString, refreshString, err := s.tokenManager.Generate(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
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

func (s *service) Logout(accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken)
	if err != nil {
		return err
	}

	id := claim.GetID()
	user, err := s.users.Find(id)
	if err != nil {
		return err
	}

	err = user.Logout()
	if err != nil {
		return err
	}

	if err := s.users.Store(user); err != nil {
		return err
	}

	return nil
}

func (s *service) CheckAndRefresh(accessToken, refreshToken string) (string, string, error) {
	accessClaim, err := s.tokenManager.Verify(accessToken)
	if accessClaim == nil {
		return "", "", err
	}

	needRefresh := err != nil
	if !needRefresh {
		return accessToken, refreshToken, nil
	}

	id := accessClaim.GetID()
	user, err := s.users.Find(id)
	if err != nil {
		return "", "", err
	}

	if user.IsBlocked {
		return "", "", authentication.ErrUserIsBlocked
	}

	if !user.IsOnline {
		return "", "", errors.New("you should login in again")
	}

	return s.tokenManager.Refresh(refreshToken)
}

func (s *service) ChangePassword(newPassword, accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken)
	if err != nil {
		return err
	}

	id := claim.GetID()
	user, err := s.users.Find(id)
	if err != nil {
		return err
	}

	err = user.ChangePassword(newPassword)
	if err != nil {
		return err
	}

	if err := s.users.Store(user); err != nil {
		return err
	}

	return nil
}

// NewService creates a auth service with necessary dependencies.
func NewService(users authentication.UserRepository, tokenManager authentication.TokenManager) Service {
	return &service{
		users: users,
		tokenManager: tokenManager,
	}
}