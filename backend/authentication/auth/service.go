package auth

import (
	"authentication"
	"errors"
	"fmt"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrEmailIsUsed is returned when email is used in register.
var ErrEmailIsUsed = errors.New("email is used")

// Service is the interface that provides authorization methods.
type Service interface {
	// Login checks user email and password, returns token.
	Login(email string, password string, ipAddr string) (string, string, error)

	// Logout.
	Logout(accessToken string) error

	// Register creates new user.
	Register(email, gender, nickname, password string) (authentication.MemberID, error)

	// Check token whether it is still valid.
	Check(accessToken string) error

	// Refresh token
	Refresh(refreshToken string) (string, string, error)

	// Resets password
	ChangePassword(newPassword, accessToken string) error
}

type service struct {
	userDB authentication.UserRepository
	userCache authentication.UserRepository
	tokenManager authentication.TokenManager
}

func (s *service) Login(email string, password string, ipAddr string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", ErrInvalidArgument
	}

	user, err := s.userDB.Find(email)
	if err != nil {
		return "", "", err
	}

	err = user.Login(password, ipAddr)
	if err != nil {
		return "", "", err
	}

	err = s.userDB.Store(user)
	if err != nil {
		return "", "", err
	}

	accessString, refreshString, err := s.tokenManager.Generate(
		user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func (s *service) Register(
		email, gender, nickname, password string) (authentication.MemberID, error) {
	if email == "" ||
	   password == "" ||
	   gender == "" ||
	   nickname == "" {
		return "", ErrInvalidArgument
	}

	if _, err := s.userDB.Find(email); err == nil {
		return "", ErrEmailIsUsed
	}

	id := authentication.NextMemberID()
	user, err := authentication.NewUser(
		id,
		"normal",
		email,
		gender,
		nickname,
		password)

	if err != nil {
		return "", err
	}

	if err := s.userDB.Store(user); err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s *service) Logout(accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken, false)
	if err != nil {
		return err
	}

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return err
	}

	claimMap := claim.ConvertToMap()
	if _, ok := claimMap["IssuedAt"]; !ok {
		return errors.New("no issued at")
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) {
		return errors.New("You should login again")
	}

	err = user.Logout()
	if err != nil {
		return err
	}

	if err := s.userDB.Store(user); err != nil {
		return err
	}

	return nil
}

func (s *service) Check(accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken, false)
	if err != nil {
		return err
	}

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return err
	}

	claimMap := claim.ConvertToMap()
	if _, ok := claimMap["IssuedAt"]; !ok {
		return errors.New("no issued at")
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) ||
	   !user.IsOnline {
		return errors.New("You should login again")
	}

	return nil
}

func (s *service) Refresh(refreshToken string) (string, string, error) {
	claim, err := s.tokenManager.Verify(refreshToken, true)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("(RefreshToken) %s", err))
	}

	claimMap := claim.ConvertToMap()

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return "", "", err
	}

	if user.IsBlocked {
		return "", "", authentication.ErrUserIsBlocked
	}

	if _, ok := claimMap["IssuedAt"]; !ok {
		return "", "", errors.New("no issued at")
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) ||
	   !user.IsOnline {
		return "", "", errors.New("You should login again")
	}

	return s.tokenManager.Generate(
		claimMap["UserID"].(authentication.MemberID),
		claimMap["Email"].(string),
		claimMap["Role"].(string))
}

func (s *service) ChangePassword(newPassword, accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken, false)
	if err != nil {
		return err
	}

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return err
	}

	claimMap := claim.ConvertToMap()
	if _, ok := claimMap["IssuedAt"]; !ok {
		return errors.New("no issued at")
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) {
		return errors.New("You should login again")
	}

	err = user.ChangePassword(newPassword)
	if err != nil {
		return err
	}

	if err := s.userDB.Store(user); err != nil {
		return err
	}

	return nil
}

// NewService creates a auth service with necessary dependencies.
func NewService(
		userDB authentication.UserRepository,
		userCache authentication.UserRepository,
		tokenManager authentication.TokenManager) Service {
	return &service{
		userDB: userDB,
		userCache: userCache,
		tokenManager: tokenManager,
	}
}