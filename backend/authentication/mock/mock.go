package mock

import (
	"authentication"
)

// UserRepository is a mock user repository.
type UserRepository struct {
	StoreFn      func(c *authentication.User) error
	StoreInvoked bool

	FindFn      func(id authentication.MemberID) (*authentication.User, error)
	FindInvoked bool

	FindAllFn      func() []*authentication.User
	FindAllInvoked bool
}

// Store calls the StoreFn.
func (r *UserRepository) Store(c *authentication.User) error {
	r.StoreInvoked = true
	return r.StoreFn(c)
}

// Find calls the FindFn.
func (r *UserRepository) Find(id authentication.MemberID) (*authentication.User, error) {
	r.FindInvoked = true
	return r.FindFn(id)
}

// FindAll calls the FindAllFn.
func (r *UserRepository) FindAll() []*authentication.User {
	r.FindAllInvoked = true
	return r.FindAllFn()
}
