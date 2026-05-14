package service

import "context"


type Service struct {
	UserService interface {
		GetUserProfile(ctx context.Context, userId string) (any, error)
	}
}

func NewService(u *UserService) *Service {
	return &Service{
		UserService: u,
	}
}


