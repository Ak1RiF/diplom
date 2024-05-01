package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/repository"
)

type PetService struct {
	petRepository repository.Pets
}

func NewPetService(repo repository.Pets) *PetService {
	return &PetService{petRepository: repo}
}

func (s *PetService) GetUserPets(userId int) ([]*dtos.OutputPet, error) {
	pets, err := s.petRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	var petsOutput []*dtos.OutputPet
	for _, v := range pets {
		petDto := dtos.OutputPet{
			Name:   v.Name,
			Rarity: v.Rarity,
		}

		petsOutput = append(petsOutput, &petDto)
	}

	return petsOutput, nil
}
func (s *PetService) GetUserPet(userId, petId int) (*dtos.OutputPet, error) {
	pet, err := s.petRepository.GetById(petId, userId)
	if err != nil {
		return nil, err
	}
	petOutput := dtos.OutputPet{
		Name:   pet.Name,
		Rarity: pet.Rarity,
	}

	return &petOutput, nil
}
func (s *PetService) AddPetToUser(userId, petId int) error {
	if err := s.petRepository.AddToUser(petId, userId); err != nil {
		return err
	}
	return nil
}
