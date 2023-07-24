package store

import (
	"GobbleGame/internal/model"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *pgx.Conn
}

func (s *Store) UserRegister(user *model.User) error {
	q := ``
	return nil
}

func (s *Store) UserExist(user *model.User) error {
	q := ``
	return nil
}

func (s *Store) UserUniqueUsername(username string) error {
	q := ``
	return nil
}

func (s *Store) UserUniqueEmail(email string) error {
	q := ``
	return nil
}
