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
}

type Quests interface {
	GetUserQuests(userId int) ([]*dtos.OutputInputDto, error)
	GetUserQuestById(questId, userId int) (*dtos.OutputInputDto, error)
	AddUserQuest(input dtos.InputQuestDto, userId int) error
	UpdateUserQuest(questId, userId int, input dtos.InputQuestDto) error
	RemoveUserQuest(questId, userId int) error
}

type Pets interface {
	GetUserPets(userId int) ([]*dtos.OutputPet, error)
	GetUserPet(userId, petId int) (*dtos.OutputPet, error)
	AddPetToUser(userId, petId int) error
}

type Eggs interface {
	GetUserEggs(userId int) ([]*dtos.OutputEgg, error)
	AddEggToUser(userId, eggId int) error
	AddToCountEgg(userId, eggId int) error
	TakeFromCountEgg(userId, eggId int) error
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
