package service

import (
	"context"

	"github.com/Santiago1010/inventory-go/internal/models"
	"github.com/Santiago1010/inventory-go/internal/repository"
)

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	ReigsterUser(ctx context.Context, email string, name string, password string) error
	LoginUser(ctx context.Context, email string, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
