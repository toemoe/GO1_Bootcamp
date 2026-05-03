package services

import (
	"s21_rogue/internal/app/handlers"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/utils"

	"github.com/go-playground/validator/v10"
)

type LoadGameService interface {
	LoadGame() bool
}

type loadGameServiceImpl struct {
	playgroundRepo repository.PlaygroundRepository
	statisticsRepo repository.StatisticsRepository
	publisher      events.EventPublisher
	log            utils.Logger
	validate       *validator.Validate
}

func NewLoadGameService(
	playgroundRepo repository.PlaygroundRepository,
	statisticsRepo repository.StatisticsRepository,
	publisher events.EventPublisher,
	validate *validator.Validate,
	log utils.Logger) LoadGameService {

	return &loadGameServiceImpl{
		playgroundRepo: playgroundRepo,
		statisticsRepo: statisticsRepo,
		publisher:      publisher,
		validate:       validate,
		log:            log,
	}

}

func (uc *loadGameServiceImpl) LoadGame() bool {
	playground, err := uc.playgroundRepo.Get(true)
	if err != nil {
		return false
	}

	stats := uc.statisticsRepo.FindLastStats()
	plStatistics, ok := stats[playground.Id]
	if !ok {
		return false
	}

	err = uc.validate.Struct(playground.Dungeon)
	if err != nil {
		return false
	}

	uc.publisher.UnsubscribeAll()
	statisticsHandler := handlers.NewStatisticsHandler(plStatistics, uc.statisticsRepo)
	uc.publisher.Subscribe(statisticsHandler,
		events.FoodEatenEvent,
		events.ElixirsDrunkEvent,
		events.ScrollsReadEvent,
		events.TreasureFoundEvent,
		events.NextLevelEvent,
		events.MonsterDeadEvent,
		events.MonsterAttacked,
		events.CharacterAttacked,
		events.CharacterStepped,
		events.LevelChangedEvent)

	messageHandler := handlers.NewMessageHandler(uc.log)

	uc.publisher.Subscribe(messageHandler,
		events.FoodEatenEvent,
		events.ElixirsDrunkEvent,
		events.ScrollsReadEvent,
		events.TreasureFoundEvent,
		events.NextLevelEvent,
		events.MonsterDeadEvent,
		events.MonsterAttacked,
		events.CharacterAttacked,
		events.CharacterStepped,
		events.LevelChangedEvent)

	uc.playgroundRepo.Save(playground, false)
	return true
}
