package handlers

import (
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/events"
)

type statisticsEventHandler struct {
	Statistics *entities.Statistics
	Repository repository.StatisticsRepository
}

func NewStatisticsHandler(statistics *entities.Statistics,
	repository repository.StatisticsRepository) events.EventHandler {

	return &statisticsEventHandler{
		Statistics: statistics,
		Repository: repository}
}

func (s *statisticsEventHandler) Handle(event events.Event) {
	switch event.GetType() {
	case events.FoodEatenEvent:
		s.Statistics.FoodEaten++
	case events.ElixirsDrunkEvent:
		s.Statistics.ElixirsDrunk++
	case events.ScrollsReadEvent:
		s.Statistics.ScrollsRead++
	case events.TreasureFoundEvent:
		s.Statistics.TreasureFound++
	case events.MonsterDeadEvent:
		s.Statistics.MonsterDead++
	case events.MonsterAttacked:
		s.Statistics.RecievedDamage++
	case events.CharacterAttacked:
		s.Statistics.GivenDamage++
	case events.CharacterStepped:
		s.Statistics.SteppedCount++
	case events.LevelChangedEvent:
		s.Statistics.Level++
		s.Repository.Save(s.Statistics)
	}
}
