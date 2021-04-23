package inmem

import (
	"auth_subsystem"
	"sync"
)

type userRepository struct {
	mtx    sync.RWMutex
	users map[string]*auth_subsystem.User
}

func (r *userRepository) Store(user *auth_subsystem.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[user.Email] = user
	return nil
}

func (r *userRepository) Find(email string) (*auth_subsystem.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.users[email]; ok {
		return val, nil
	}
	return nil, auth_subsystem.ErrUnknownUser
}

func (r *userRepository) FindAll() []*auth_subsystem.User {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	users := make([]*auth_subsystem.User, 0, len(r.users))
	for _, val := range r.users {
		users = append(users, val)
	}
	return users
}

// NewUserRepository returns a new instance of a in-memory user repository.
func NewUserRepository () auth_subsystem.UserRepository {
	return &userRepository{
		users: make(map[string]*auth_subsystem.User),
	}
}