package redis

// import (
// 	"context"
// 	"sync"

// 	"github.com/go-redis/redis/v8"
// )

// type userRepository struct {
// 	ctx    context.Context
// 	client *redis.Client
// }

// func (r *userRepository) Store(user *authentication.User) error {
// 	r.client.Set(r.ctx, )
// 	r.users[user.ID] = user
// 	return nil
// }

// func (r *userRepository) Find(id authentication.MemberID) (*authentication.User, error) {
// 	r.mtx.Lock()
// 	defer r.mtx.Unlock()
// 	if val, ok := r.users[id]; ok {
// 		return val, nil
// 	}
// 	return nil, authentication.ErrUnknownUser
// }

// func (r *userRepository) FindAll() []*authentication.User {
// 	r.mtx.Lock()
// 	defer r.mtx.Unlock()
// 	users := make([]*authentication.User, 0, len(r.users))
// 	for _, val := range r.users {
// 		users = append(users, val)
// 	}
// 	return users
// }