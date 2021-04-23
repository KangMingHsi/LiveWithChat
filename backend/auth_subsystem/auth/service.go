package auth

import (
	"auth_subsystem"
	"errors"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrEmailIsUsed is returned when email is used in register.
var ErrEmailIsUsed = errors.New("email is used")

// Service is the interface that provides authorization methods.
type Service interface {
	// Login checks user email and password, returns token.
	Login(email string, password string, ipAddr string) (string, error)

	// Logout.
	Logout(accessToken string) error

	// Register creates new user.
	Register(email, gender, nickname, password string) (auth_subsystem.MemberID, error)

	// Check token whether it is still valid.
	Check(accessToken string) (auth_subsystem.MemberID, int, error)

	// Refresh token
	Refresh(refreshToken string) (string, error)

	// Resets password
	ChangePassword(newPassword, accessToken string) error
}

type service struct {
	userDB auth_subsystem.UserRepository
	userCache auth_subsystem.UserRepository
	tokenManager auth_subsystem.TokenManager
}

func (s *service) Login(email string, password string, ipAddr string) (string, error) {
	if email == "" || password == "" {
		return "", ErrInvalidArgument
	}

	user, err := s.userDB.Find(email)
	if err != nil {
		return "", err
	}

	err = user.Login(password, ipAddr)
	if err != nil {
		return "", err
	}

	err = s.userDB.Store(user)
	if err != nil {
		return "", err
	}

	accessString, err := s.tokenManager.Generate(
		string(user.ID), user.Email, user.RoleLevel())
	if err != nil {
		return "", err
	}

	return accessString, nil
}

func (s *service) Register(
		email, gender, nickname, password string) (auth_subsystem.MemberID, error) {
	if email == "" ||
	   password == "" ||
	   gender == "" ||
	   nickname == "" {
		return "", ErrInvalidArgument
	}

	if _, err := s.userDB.Find(email); err == nil {
		return "", ErrEmailIsUsed
	}

	id := auth_subsystem.NextMemberID()
	user, err := auth_subsystem.NewUser(
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
	claim, err := s.tokenManager.Verify(accessToken)
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
		return auth_subsystem.ErrNoIssuedAt
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) {
		return auth_subsystem.ErrUserShouldLogin
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

func (s *service) Check(accessToken string) (auth_subsystem.MemberID, int, error) {
	claim, err := s.tokenManager.Verify(accessToken)
	if err != nil {
		return "", 0, err
	}

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return "", 0, err
	}

	claimMap := claim.ConvertToMap()
	if _, ok := claimMap["IssuedAt"]; !ok {
		return "", 0, auth_subsystem.ErrNoIssuedAt
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) ||
	   !user.IsOnline {
		return "", 0, auth_subsystem.ErrUserShouldLogin
	}

	return user.ID, user.RoleLevel(), nil
}

func (s *service) Refresh(refreshToken string) (string, error) {
	claim, err := s.tokenManager.VerifyWithoutExpired(refreshToken)
	if err != nil {
		return "", err
	}

	claimMap := claim.ConvertToMap()

	email := claim.GetKey().(string)
	user, err := s.userDB.Find(email)
	if err != nil {
		return "", err
	}

	if user.IsBlocked {
		return "", auth_subsystem.ErrUserIsBlocked
	}

	if _, ok := claimMap["IssuedAt"]; !ok {
		return "", auth_subsystem.ErrNoIssuedAt
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) ||
	   !user.IsOnline {
		return "", auth_subsystem.ErrUserShouldLogin
	}

	return s.tokenManager.Generate(
		claimMap["UserID"].(string),
		claimMap["Email"].(string),
		claimMap["RoleLevel"].(int))
}

func (s *service) ChangePassword(newPassword, accessToken string) error {
	claim, err := s.tokenManager.Verify(accessToken)
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
		return auth_subsystem.ErrNoIssuedAt
	}

	if user.LimitationPeriod.Unix() > claimMap["IssuedAt"].(int64) {
		return auth_subsystem.ErrUserShouldLogin
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
		userDB auth_subsystem.UserRepository,
		userCache auth_subsystem.UserRepository,
		tokenManager auth_subsystem.TokenManager) Service {
	return &service{
		userDB: userDB,
		userCache: userCache,
		tokenManager: tokenManager,
	}
}