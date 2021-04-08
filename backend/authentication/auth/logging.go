package auth

import (
	"time"

	"github.com/go-kit/kit/log"

	"authentication"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Login(email string, password string, ipAddr string) (accessToken string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Login(email, password, ipAddr)
}

func (s *loggingService) Logout(accessToken string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Logout",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Logout(accessToken)
}

func (s *loggingService) Register(email, gender, nickname, password string) (id authentication.MemberID, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Register",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Register(email, gender, nickname, password)
}

func (s *loggingService) Check(accessToken string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Check",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Check(accessToken)
}

func (s *loggingService) Refresh(refreshToken string) (newAccessToken string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Refresh",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Refresh(refreshToken)
}

func (s *loggingService) ChangePassword(newPassword, accessToken string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ChangePassword",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.ChangePassword(newPassword, accessToken)
}