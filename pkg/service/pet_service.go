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
		name, _ := s.GetNameOfPet(userId, v.Id)
		petDto := dtos.OutputPet{
			Id:     v.Id,
			Name:   name,
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
	name, _ := s.GetNameOfPet(userId, petId)
	petOutput := dtos.OutputPet{
		Id:     pet.Id,
		Name:   name,
		Rarity: pet.Rarity,
	}

	return &petOutput, nil
}
func (s *PetService) AddPetToUser(userId, petId int, name string) error {
	if err := s.petRepository.AddToUser(petId, userId, name); err != nil {
		return err
	}
	return nil
}

func (s *PetService) ChangePetName(userId, petId int, name string) error {
	if err := s.petRepository.ChangeName(petId, userId, name); err != nil {
		return err
	}
	return nil
}

func (s *PetService) GetNameOfPet(userId, petId int) (string, error) {
	name, err := s.petRepository.GetNamePet(petId, userId)
	if err != nil {
		return "", err
	}
	return name, nil
}
