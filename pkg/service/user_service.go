package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/repository"
)

type UserService struct {
	userRepository repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) UserInfo(id int) (*dtos.OutputUserDto, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dtos.OutputUserDto{
		Username:              user.Username,
		TotalExperience:       user.TotalExperience,
		AmountExperienceToLvl: user.AmountExperienceToLvl,
		Lvl:                   user.Lvl,
	}, nil
}
