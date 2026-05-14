package store

import "context"

type Storage struct {
	User interface {
		CreateUser() error
		FetchByID(ctx context.Context, id string) (any, error)
	}
}

func NewStorage(u *UserStruct) *Storage {
	return &Storage{
		User: u,
	}
}
