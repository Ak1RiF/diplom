package service

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/helpers"
	"github.com/todoApp/pkg/models"
	"github.com/todoApp/pkg/repository"
)

type AuthorizationService struct {
	userRepository repository.Users
}

func NewAuthService(repo repository.Users) *AuthorizationService {
	return &AuthorizationService{
		userRepository: repo,
	}
}

// implementations
func (s *AuthorizationService) AddUser(input dtos.InputUserForm) error {
	passwordHash, err := helpers.GeneratePasswordHash(input.Password)
	if err != nil {
		return err
	}
	user := models.User{
		Username:     input.Username,
		PasswordHash: passwordHash,
	}

	if err := s.userRepository.Create(user); err != nil {
		//
		return err
	}
	return err
}

func (s *AuthorizationService) GenerateJwt(input dtos.InputUserForm) (string, error) {
	user, err := s.userRepository.GetByUsername(input.Username)
	if err != nil {
		return "", err
	}

	if !helpers.VerifyPasswordHash(user.PasswordHash, input.Password) {
		return "Wrong login or password", errors.New("wrong input data")
	}

	token, err := helpers.GenerateToken(*user)
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *AuthorizationService) ParseJwt(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &helpers.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return helpers.SigningKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*helpers.CustomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type (*CustomClaims)")
	}

	return claims.UserId, nil
}
