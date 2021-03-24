package auth

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"authentication"
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

func (s *instrumentingService) Login(id authentication.MemberID, password string) (token string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Login").Add(1)
		s.requestLatency.With("method", "Login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Login(id, password)
}

func (s *instrumentingService) Logout(id authentication.MemberID) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Logout").Add(1)
		s.requestLatency.With("method", "Logout").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Logout(id)
}

func (s *instrumentingService) Register(password string) (id authentication.MemberID, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Register").Add(1)
		s.requestLatency.With("method", "Register").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Register(password)
}
