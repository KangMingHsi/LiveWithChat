package auth

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"auth_subsystem"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) Login(email string, password string, ipAddr string) (
		loginInfo map[string]string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Login").Add(1)
		s.requestLatency.With("method", "Login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Login(email, password, ipAddr)
}

func (s *instrumentingService) Logout(accessToken string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Logout").Add(1)
		s.requestLatency.With("method", "Logout").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Logout(accessToken)
}

func (s *instrumentingService) Register(email, gender, nickname, password string) (id auth_subsystem.MemberID, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Register").Add(1)
		s.requestLatency.With("method", "Register").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Register(email, gender, nickname, password)
}

func (s *instrumentingService) Check(accessToken string) (uid auth_subsystem.MemberID, level int, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Check").Add(1)
		s.requestLatency.With("method", "Check").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Check(accessToken)
}

func (s *instrumentingService) Refresh(refreshToken string) (newAccessToken string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Refresh").Add(1)
		s.requestLatency.With("method", "Refresh").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Refresh(refreshToken)
}

func (s *instrumentingService) ChangePassword(newPassword, accessToken string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CheckOrRefresh").Add(1)
		s.requestLatency.With("method", "CheckOrRefresh").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.ChangePassword(newPassword, accessToken)
}