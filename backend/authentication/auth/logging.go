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

func (s *loggingService) Login(id authentication.MemberID, password string) (token string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Login(id, password)
}

func (s *loggingService) Logout(id authentication.MemberID) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Logout(id)
}

func (s *loggingService) Register(password string) (id authentication.MemberID, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.Register(password)
}