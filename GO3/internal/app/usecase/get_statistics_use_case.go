package usecase

import (
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
)

type GetStatisticsUseCase interface {
	GetStatistics() []*dto.StatisticsDTO
}

type getStatisticsUseCaseImpl struct {
	repo repository.StatisticsRepository
	log  utils.Logger
}

func NewGetStatisticsUseCase(repo repository.StatisticsRepository, log utils.Logger) GetStatisticsUseCase {
	return &getStatisticsUseCaseImpl{
		repo: repo,
		log:  log}

}

func (uc *getStatisticsUseCaseImpl) GetStatistics() []*dto.StatisticsDTO {
	leaderStatistics := uc.repo.FindLeaderOrderByTreasure(10)

	leaderDTO := make([]*dto.StatisticsDTO, 0)
	for _, v := range leaderStatistics {
		leaderDTO = append(leaderDTO, mappers.MapToStatisticsDTO(v))
	}
	return leaderDTO
}
