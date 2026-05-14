package service

import (
	"context"

	"github.com/hanzala211/go-backend-template/internal/store"
	"golang.org/x/sync/singleflight"
)

type UserService struct {
	store        *store.Storage
	requestGroup singleflight.Group
}

func NewUserService(s *store.Storage) *UserService {
	return &UserService{
		store: s,
	}
}

func (u *UserService) GetUserProfile(ctx context.Context, userId string) (any, error) {
	cacheKey := "user-" + userId

	// cache check goes here
	v, err, _ := u.requestGroup.Do(cacheKey, func() (any, error) {
		user, err := u.store.User.FetchByID(ctx, userId)
		if err != nil {
			return nil, err
		}
		// cache add logic will go here
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	user := v.(store.User)
	return user, err
}
