package store

import "database/sql"

type UserStruct struct {
	db *sql.DB
}

func NewUserStruct(db *sql.DB) *UserStruct {
	return &UserStruct{db: db}
}

func (u *UserStruct) CreateUser() error {
	return nil
}
