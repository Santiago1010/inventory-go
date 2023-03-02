package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Santiago1010/inventory-go/encryption"
	"github.com/Santiago1010/inventory-go/internal/models"
)

var (
	ErrorUserAlreadyExists  = errors.New("el usuario existe")
	ErrorInvalidCredentials = errors.New("credenciales inválidas")
)

func (s *serv) ReigsterUser(ctx context.Context, email string, name string, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)

	if u != nil {
		return ErrorUserAlreadyExists
	}

	encryptedPassword, err := encryption.HashPassword(password)

	if err != nil {
		return ErrorInvalidCredentials
	}

	fmt.Println(string(encryptedPassword))

	return s.repo.SaveUser(ctx, email, name, string(encryptedPassword))
}

func (s *serv) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if u.Password != password {
		return nil, ErrorInvalidCredentials
	}

	return &models.User{
		Id:    u.Id,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
