package repository

import (
	"github.com/todoApp/pkg/helpers"
	"github.com/todoApp/pkg/models"
)

// interfaces
type Users interface {
	GetByUsername(username string) (*models.User, error)
	Create(user models.User) error
}

type Quests interface {
}

type Pets interface {
}

type Eggs interface {
}

// repository struct
type Repository struct {
	Users
}

func NewRepository() *Repository {
	err := helpers.Init()
	if err != nil {
		//log
		return nil
	}

	databaseUrl := helpers.GetByKey("DATABASE_URL")
	return &Repository{
		Users: NewUserRepository(databaseUrl),
	}
}
