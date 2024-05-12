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

func (s *EggService) GetUserEggs(userId int) (*dtos.OutputEggs, error) {
	// eggs, err := s.eggRepository.Get(userId)
	// if err != nil {
	// 	return nil, err
	// }
	// var eggsOutput []*dtos.OutputEgg
	// for _, v := range eggs {
	// 	eggDto := dtos.OutputEgg{
	// 		Rarity: v.Rarity,
	// 	}

	// 	eggsOutput = append(eggsOutput, &eggDto)
	// }

	// return eggsOutput, nil
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
