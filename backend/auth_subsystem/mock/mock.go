package mock

import (
	"auth_subsystem"
)

// UserRepository is a mock user repository.
type UserRepository struct {
	StoreFn      func(c *auth_subsystem.User) error
	StoreInvoked bool

	FindFn      func(id auth_subsystem.MemberID) (*auth_subsystem.User, error)
	FindInvoked bool

	FindAllFn      func() []*auth_subsystem.User
	FindAllInvoked bool
}

// Store calls the StoreFn.
func (r *UserRepository) Store(c *auth_subsystem.User) error {
	r.StoreInvoked = true
	return r.StoreFn(c)
}

// Find calls the FindFn.
func (r *UserRepository) Find(id auth_subsystem.MemberID) (*auth_subsystem.User, error) {
	r.FindInvoked = true
	return r.FindFn(id)
}

// FindAll calls the FindAllFn.
func (r *UserRepository) FindAll() []*auth_subsystem.User {
	r.FindAllInvoked = true
	return r.FindAllFn()
}
