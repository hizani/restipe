package postgres

import (
	"fmt"
	"restipe/internal/model"

	"github.com/jmoiron/sqlx"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStoarge(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{db}
}

func (s *AuthStorage) SigninUser(user model.SigninUserReq) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err := s.db.Get(&id, query, user.Login, user.Password)

	return id, err
}

func (s *AuthStorage) SignupUser(user model.SignupUserReq) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := s.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AuthStorage) Close() error {
	return s.db.Close()
}
