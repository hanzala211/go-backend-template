package store

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserStruct struct {
	db *sql.DB
}

func NewUserStruct(db *sql.DB) *UserStruct {
	return &UserStruct{db: db}
}

func (u *UserStruct) CreateUser() error {
	return nil
}

func (u *UserStruct) FetchByID(ctx context.Context, id string) (any, error) {
	query := `SELECT id, name FROM users u WHERE id = $1;`
	var user User
	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User not found!")
		}
		return nil, err
	}
	return user, nil
}
