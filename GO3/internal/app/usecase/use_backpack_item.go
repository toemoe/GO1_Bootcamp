package usecase

import (
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/utils"

	"github.com/google/uuid"
)

type UseBackpackUseCase interface {
	UseItem(id string)
}

type useBackpackUseCaseImpl struct {
	repo      repository.PlaygroundRepository
	publisher events.EventPublisher
	log       utils.Logger
}

func NewUseBackpackUseCase(repo repository.PlaygroundRepository,
	publisher events.EventPublisher,
	log utils.Logger) UseBackpackUseCase {

	return &useBackpackUseCaseImpl{repo: repo,
		publisher: publisher,
		log:       log}
}

func (c *useBackpackUseCaseImpl) UseItem(id string) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.log.Error("Incorrect parse uuid ")
		return
	}
	playground, _ := c.repo.Get(false)
	defer c.repo.Save(playground, false)

	event := playground.UseItemByUUID(uuid)
	c.publisher.Notify(event)
}
