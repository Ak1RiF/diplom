package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/repository"
)

type EggService struct {
	eggRepository repository.Eggs
}

func NewEggService(repo repository.Eggs) *EggService {
	return &EggService{
		eggRepository: repo,
	}
}

func (s *EggService) GetUserEggs(userId int) ([]*dtos.OutputEgg, error) {
	eggs, err := s.eggRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	var eggsOutput []*dtos.OutputEgg
	for _, v := range eggs {
		eggDto := dtos.OutputEgg{
			Rarity: v.Rarity,
			Count:  v.Count,
		}

		eggsOutput = append(eggsOutput, &eggDto)
	}

	return eggsOutput, nil
}
func (s *EggService) AddEggToUser(userId, eggId int) error {
	if err := s.eggRepository.AddToUser(eggId, userId); err != nil {
		return err
	}
	return nil
}
