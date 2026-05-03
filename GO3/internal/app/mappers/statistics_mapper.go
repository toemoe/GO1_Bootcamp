package mappers

import (
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"

	"github.com/google/uuid"
)

func MapToStatisticsDTO(statistics *entities.Statistics) *dto.StatisticsDTO {
	statisticsDTO := dto.StatisticsDTO{}
	statisticsDTO.Id = statistics.Id.String()
	statisticsDTO.Level = statistics.Level

	statisticsDTO.MonsterDead = statistics.MonsterDead
	statisticsDTO.GivenDamage = statistics.GivenDamage
	statisticsDTO.RecievedDamage = statistics.GivenDamage

	statisticsDTO.FoodEaten = statistics.FoodEaten
	statisticsDTO.ElixirsDrunk = statistics.ElixirsDrunk
	statisticsDTO.ScrollsRead = statistics.ScrollsRead
	statisticsDTO.TreasureFound = statistics.TreasureFound

	statisticsDTO.SteppedCount = statistics.SteppedCount

	return &statisticsDTO
}

func MapFromStatisticsDTO(statisticsDTO *dto.StatisticsDTO) (*entities.Statistics, error) {
	id, err := uuid.Parse(statisticsDTO.Id)
	if err != nil {
		return nil, err
	}

	s := entities.NewStatistics(id)
	s.Level = statisticsDTO.Level

	s.MonsterDead = statisticsDTO.MonsterDead
	s.GivenDamage = statisticsDTO.GivenDamage
	s.RecievedDamage = statisticsDTO.RecievedDamage

	s.FoodEaten = statisticsDTO.FoodEaten
	s.ElixirsDrunk = statisticsDTO.ElixirsDrunk
	s.ScrollsRead = statisticsDTO.ScrollsRead
	s.TreasureFound = statisticsDTO.TreasureFound

	s.SteppedCount = statisticsDTO.SteppedCount

	return s, nil
}
