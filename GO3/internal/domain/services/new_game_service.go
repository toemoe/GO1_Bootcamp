package services

import (
	"s21_rogue/internal/app/handlers"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/utils"
)

type NewGameService interface {
	NewGame()
}

type newGameServiceImpl struct {
	playgroundGeneratorService PlaygroundGeneratorService
	playgroundRepo             repository.PlaygroundRepository
	statisticsRepo             repository.StatisticsRepository
	publisher                  events.EventPublisher
	log                        utils.Logger
}

func NewNewGameService(
	playgroundGeneratorService PlaygroundGeneratorService,
	playgroundRepo repository.PlaygroundRepository,
	statisticsRepo repository.StatisticsRepository,
	publisher events.EventPublisher,
	log utils.Logger) NewGameService {

	return &newGameServiceImpl{
		playgroundGeneratorService: playgroundGeneratorService,
		playgroundRepo:             playgroundRepo,
		statisticsRepo:             statisticsRepo,
		publisher:                  publisher,
		log:                        log,
	}

}

func (uc *newGameServiceImpl) NewGame() {
	uc.publisher.UnsubscribeAll()

	playground := uc.playgroundGeneratorService.GenerateNewLevel(nil, 1)
	uc.playgroundRepo.Save(playground, false)

	statistics := entities.NewStatistics(playground.Id)
	statisticsHandler := handlers.NewStatisticsHandler(statistics, uc.statisticsRepo)
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

}
