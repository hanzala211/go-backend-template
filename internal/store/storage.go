package store

type Storage struct {
	User interface {
		CreateUser() error
	}
}

func NewStorage(u *UserStruct) *Storage {
	return &Storage{
		User: u,
	}
}
