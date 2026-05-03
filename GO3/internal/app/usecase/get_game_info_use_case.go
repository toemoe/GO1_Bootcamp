package usecase

import (
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"

	"github.com/google/uuid"
)

type GetGameInfoUseCase interface {
	GetGameInfo(id uuid.UUID) *dto.GameInfoDTO
}

type getGameInfoUseCaseImpl struct {
	gcRepo    repository.GameStateRepository
	levelRepo repository.StatisticsRepository
	log       utils.Logger
}

func NewGetGameInfoUseCase(gcRepo repository.GameStateRepository,
	levelRepo repository.StatisticsRepository,
	log utils.Logger) GetGameInfoUseCase {

	return &getGameInfoUseCaseImpl{
		gcRepo:    gcRepo,
		levelRepo: levelRepo,
		log:       log}
}

func (c *getGameInfoUseCaseImpl) GetGameInfo(id uuid.UUID) *dto.GameInfoDTO {
	gameInfo := c.gcRepo.Get()
	level := c.levelRepo.GetLevelById(id)

	gameInfoDTO := mappers.MapToGameInfoDTO(gameInfo, level)
	return gameInfoDTO
}
