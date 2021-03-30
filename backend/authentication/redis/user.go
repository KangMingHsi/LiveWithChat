package redis

import (
	"authentication"
	"context"

	"github.com/go-redis/cache/v8"
)

type userRepository struct {
	prefix string
	ctx	context.Context
	uCache *cache.Cache
}

func (r *userRepository) Store(user *authentication.User) error {
	return r.uCache.Set(
		&cache.Item{
			Ctx: r.ctx,
			Key: r.prefix + string(user.ID),
			Value: user,
		},
	)
}

func (r *userRepository) Find(id authentication.MemberID) (*authentication.User, error) {
	var user *authentication.User
	err := r.uCache.Get(r.ctx, r.prefix + string(id), &user)
	if err != nil {
		return nil, authentication.ErrUnknownUser
	}
	return user, nil
}

func (r *userRepository) FindAll() []*authentication.User {
	return []*authentication.User{}
}

func NewUserRepository (client *cache.Cache) authentication.UserRepository {
	return &userRepository{
		prefix: "authUser_",
		uCache: client,
		ctx: context.Background(),
	}
}