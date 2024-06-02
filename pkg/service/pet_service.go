package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/models"
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
			Id:     v.Id,
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
		Id:     pet.Id,
		Name:   pet.Name,
		Rarity: pet.Rarity,
	}

	return &petOutput, nil
}

func (s *PetService) CreatePetToUser(userId int, input dtos.CreatePet) error {
	if err := s.petRepository.Create(userId, models.Pet{Name: input.Name, Rarity: input.Rarity}); err != nil {
		return err
	}
	return nil
}

func (s *PetService) ChangePet(userId int, input dtos.UpdatePet) error {
	pet, err := s.petRepository.GetById(input.Id, userId)
	if err != nil {
		return err
	}

	pet.Name = input.Name
	if err := s.petRepository.Update(pet.Id, userId, *pet); err != nil {
		return err
	}
	return nil
}
