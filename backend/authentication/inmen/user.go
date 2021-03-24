package inmen

import (
	"authentication"
	"sync"
)

type userRepository struct {
	mtx    sync.RWMutex
	users map[authentication.MemberID]*authentication.User
}

func (r *userRepository) Store(user *authentication.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) Find(id authentication.MemberID) (*authentication.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.users[id]; ok {
		return val, nil
	}
	return nil, authentication.ErrUnknownUser
}

func (r *userRepository) FindAll() []*authentication.User {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	users := make([]*authentication.User, 0, len(r.users))
	for _, val := range r.users {
		users = append(users, val)
	}
	return users
}

// NewUserRepository returns a new instance of a in-memory user repository.
func NewUserRepository () authentication.UserRepository {
	return &userRepository{
		users: make(map[authentication.MemberID]*authentication.User),
	}
}