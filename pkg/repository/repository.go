package repository

import (
	"github.com/todoApp/pkg/helpers"
	"github.com/todoApp/pkg/models"
)

// interfaces
type Users interface {
	GetById(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Create(user models.User) error
	Update(userId int, user models.User) error
	UpdateExperience(userId int, user models.User) error
}

type Quests interface {
	Get(userId int) ([]models.Quest, error)
	GetById(id, userId int) (*models.Quest, error)
	Create(quest models.Quest, userId int) (int, error)
	Update(id, userId int, quest models.Quest) error
	Delete(id, userId int) error

	//	GetCompletedQuests(userId int) error
	PointAsCompletedById(id, userId int) error
}

type Pets interface {
	Get(userId int) ([]models.Pet, error)
	GetById(petId, userId int) (*models.Pet, error)
	GetNamePet(petId, userId int) (string, error)
	AddToUser(petId, userId int, name string) error
	ChangeName(petId, userId int, name string) error
}

type Eggs interface {
	Get(userId int) ([]int, error)
	AddToUser(eggId, userId int) error
	UpdateCount(count, eggId, userId int) error
}

// repository struct
type Repository struct {
	Users
	Quests
	Pets
	Eggs
}

func NewRepository() *Repository {
	err := helpers.Init()
	if err != nil {
		//log
		return nil
	}

	databaseUrl := helpers.GetByKey("DATABASE_URL")
	return &Repository{
		Users:  NewUserRepository(databaseUrl),
		Quests: NewQuestRepository(databaseUrl),
		Pets:   NewPetRepository(databaseUrl),
		Eggs:   NewEggRepository(databaseUrl),
	}
}
