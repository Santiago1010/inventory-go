package service

import (
	"context"
	"testing"

	"github.com/Santiago1010/inventory-go/internal/entity"
	"github.com/Santiago1010/inventory-go/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "scorrea44@protonmail.com",
			UserName:      "Test",
			Password:      "ValidPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserExist",
			Email:         "scorrea44@exist.com",
			UserName:      "Test",
			Password:      "ValidPassword",
			ExpectedError: ErrorUserAlreadyExists,
		},
	}

	ctx := context.Background()

	repo := &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exist.com").Return(&entity.User{Email: "test@exist.com"}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			s := New(repo)

			err := s.ReigsterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != nil {
				t.Errorf("Error esperado %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
