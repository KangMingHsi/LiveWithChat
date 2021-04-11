package redis

import (
	"auth_subsystem"
	"context"

	"github.com/go-redis/cache/v8"
)

type userRepository struct {
	prefix string
	ctx	context.Context
	uCache *cache.Cache
}

func (r *userRepository) Store(user *auth_subsystem.User) error {
	return r.uCache.Set(
		&cache.Item{
			Ctx: r.ctx,
			Key: r.prefix + string(user.ID),
			Value: user,
		},
	)
}

func (r *userRepository) Find(email string) (*auth_subsystem.User, error) {
	var user *auth_subsystem.User
	err := r.uCache.Get(r.ctx, r.prefix + email, &user)
	if err != nil {
		return nil, auth_subsystem.ErrUnknownUser
	}
	return user, nil
}

func (r *userRepository) FindAll() []*auth_subsystem.User {
	return []*auth_subsystem.User{}
}

func NewUserRepository (client *cache.Cache) auth_subsystem.UserRepository {
	return &userRepository{
		prefix: "authUser_",
		uCache: client,
		ctx: context.Background(),
	}
}