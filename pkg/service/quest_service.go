package service

import (
	"github.com/todoApp/pkg/dtos"
	"github.com/todoApp/pkg/models"
	"github.com/todoApp/pkg/repository"
)

type QuestService struct {
	questRepository repository.Quests
}

func NewQuestService(repo repository.Quests) *QuestService {
	return &QuestService{questRepository: repo}
}

func (s *QuestService) GetUserQuests(userId int) ([]*dtos.OutputInputDto, error) {
	quests, err := s.questRepository.Get(userId)
	if err != nil {
		// log
		return nil, err
	}

	var questsOutput []*dtos.OutputInputDto
	for _, v := range quests {
		questDto := dtos.OutputInputDto{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Dificulty:   v.Dificulty,
			Completed:   v.Completed,
		}

		questsOutput = append(questsOutput, &questDto)
	}

	return questsOutput, nil
}
func (s *QuestService) GetUserQuestById(questId, userId int) (*dtos.OutputInputDto, error) {
	quest, err := s.questRepository.GetById(questId, userId)
	if err != nil {
		// log
		return nil, err
	}
	questOutput := dtos.OutputInputDto{
		Id:          quest.Id,
		Title:       quest.Title,
		Description: quest.Description,
		Dificulty:   quest.Dificulty,
		Completed:   quest.Completed,
	}

	return &questOutput, nil
}
func (s *QuestService) AddUserQuest(input dtos.InputQuestDto, userId int) error {
	quest := models.Quest{
		Title:       input.Title,
		Description: input.Description,
		Dificulty:   input.Dificulty,
		Completed:   false,
	}

	_, err := s.questRepository.Create(quest, userId)
	if err != nil {
		// log
		return err
	}
	// log
	return nil
}
func (s *QuestService) UpdateUserQuest(questId, userId int, input dtos.InputQuestDto) error {
	quest := models.Quest{
		Title:       input.Title,
		Description: input.Description,
		Dificulty:   input.Dificulty,
	}

	if err := s.questRepository.Update(questId, userId, quest); err != nil {
		// log
		return err
	}
	// log
	return nil
}
func (s *QuestService) RemoveUserQuest(questId, userId int) error {
	if err := s.questRepository.Delete(questId, userId); err != nil {
		// log
		return err
	}
	// log
	return nil
}

func (s *QuestService) CompleteQuest(questId, userId int) error {
	if err := s.questRepository.PointAsCompletedById(questId, userId); err != nil {
		return err
	}
	return nil
}
