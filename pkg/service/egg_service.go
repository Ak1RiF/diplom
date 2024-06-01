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
func (s *EggService) GetUserEggById(userId, eggId int) (*dtos.OutputEgg, error) {
	counts, err := s.eggRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	switch eggId - 1 {
	case 0:
		return &dtos.OutputEgg{Id: 1, Rarity: "common", Count: counts[0]}, nil
	case 1:
		return &dtos.OutputEgg{Id: 2, Rarity: "rare", Count: counts[1]}, nil
	case 2:
		return &dtos.OutputEgg{Id: 3, Rarity: "epic", Count: counts[2]}, nil
	case 3:
		return &dtos.OutputEgg{Id: 4, Rarity: "legendary", Count: counts[3]}, nil
	default:
		return nil, err
	}
}

func (s *EggService) GetUserEggs(userId int) (*dtos.OutputEggs, error) {
	counts, err := s.eggRepository.Get(userId)
	if err != nil {
		return nil, err
	}

	countEggs := dtos.OutputEggs{
		CountCommon:    counts[0],
		CountRare:      counts[1],
		CountEpic:      counts[2],
		CountLegendary: counts[3],
	}

	return &countEggs, nil
}

func (s *EggService) AddEggToUser(userId, eggId int) error {
	if err := s.eggRepository.AddToUser(eggId, userId); err != nil {
		return err
	}
	return nil
}

func (s *EggService) UpdateCountEggs(count, eggId, userId int) error {
	if err := s.eggRepository.UpdateCount(count, eggId, userId); err != nil {
		return err
	}
	return nil
}
