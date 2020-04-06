package database

import (
	"database/sql"

	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type User interface {
	Add(user entity.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (dc *UserRepository) Add(user entity.User) error {

	sql := "INSERT INTO user (name, email, description) VALUES (?, ?, ?)"
	_, err := dc.db.Exec(sql, user.Name, user.Email, user.Description)
	if err != nil {
		return err
	}

	return nil
}
