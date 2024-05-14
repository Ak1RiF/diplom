package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/models"
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

func (s *UserService) UpdateUserExperience(userId int, input dtos.UserExperienceInput) error {
	user := models.User{
		TotalExperience:       input.AddToCount,
		AmountExperienceToLvl: input.AmountToLvl,
	}

	err := s.userRepository.UpdateExperience(userId, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(userId int, updateUser dtos.UpdateUserFrom) error {
	user := models.User{
		Username:              updateUser.Username,
		TotalExperience:       updateUser.TotalExperience,
		AmountExperienceToLvl: updateUser.AmountExperienceToLvl,
		Lvl:                   updateUser.Lvl,
	}

	if err := s.userRepository.Update(userId, user); err != nil {
		return err
	}
	return nil
}
