package repository

import (
	"context"

	"github.com/Santiago1010/inventory-go/internal/entity"
)

const (
	queryInsertUser  = `INSERT INTO users (email, name, password) VALUES (:email, :name, :password)`
	queryUserByEmail = `SELECT * FROM users WHERE email = :email`
)

func (r *repo) SaveUser(ctx context.Context, email string, name string, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)

	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	err := r.db.GetContext(ctx, u, queryUserByEmail, email)

	return u, err
}
