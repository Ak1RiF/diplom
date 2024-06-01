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

type Users interface {
	UserInfo(id int) (*dtos.OutputUserDto, error)
	UpdateUser(userId int, updateUser dtos.UpdateUserFrom) error
	UpdateUserExperience(userId int, input dtos.UserExperienceInput) error
}

type Quests interface {
	GetUserQuests(userId int) ([]*dtos.OutputInputDto, error)
	GetUserQuestById(questId, userId int) (*dtos.OutputInputDto, error)
	AddUserQuest(input dtos.InputQuestDto, userId int) error
	UpdateUserQuest(questId, userId int, input dtos.InputQuestDto) error
	RemoveUserQuest(questId, userId int) error
	CompleteQuest(questId, userId int) error
}

type Pets interface {
	GetUserPets(userId int) ([]*dtos.OutputPet, error)
	GetUserPet(userId, petId int) (*dtos.OutputPet, error)
	AddPetToUser(userId, petId int, name string) error
	ChangePetName(userId, petId int, name string) error
}

type Eggs interface {
	GetUserEggById(userId, eggId int) (*dtos.OutputEgg, error)
	GetUserEggs(userId int) (*dtos.OutputEggs, error)
	AddEggToUser(userId, eggId int) error
	UpdateCountEggs(count, eggId, userId int) error
}

// Service struct
type Service struct {
	Authorization
	Quests
	Pets
	Eggs
	Users
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Users),
		Quests:        NewQuestService(repository.Quests),
		Pets:          NewPetService(repository.Pets),
		Eggs:          NewEggService(repository.Eggs),
		Users:         NewUserService(repository.Users),
	}
}
