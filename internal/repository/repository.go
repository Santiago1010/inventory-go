package repository

import (
	"context"

	"github.com/Santiago1010/inventory-go/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email string, name string, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
