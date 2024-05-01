package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/repository"
)

// interfaces
type Authorization interface {
	AddUser(input dtos.InputUserForm) error
	GenerateJwt(input dtos.InputUserForm) (string, error)
	ParseJwt(accessToken string) (int, error)
}

// Service struct
type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Users),
	}
}
