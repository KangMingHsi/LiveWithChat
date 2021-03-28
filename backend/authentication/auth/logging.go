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

func (s *loggingService) Login(id authentication.MemberID, password string, ipAddr string) (accessToken, refreshToken string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Login(id, password, ipAddr)
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

func (s *loggingService) Register(password string) (id authentication.MemberID, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Register",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Register(password)
}

func (s *loggingService) CheckAndRefresh(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CheckOrRefresh",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.CheckAndRefresh(accessToken, refreshToken)
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