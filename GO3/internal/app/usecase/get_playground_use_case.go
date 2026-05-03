package usecase

import (
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
)

type GetPlaygroundUseCase interface {
	GetPlayground() *dto.PlaygroundDTO
}

type getPlaygroundUseCaseImpl struct {
	repo repository.PlaygroundRepository
	log  utils.Logger
}

func NewGetPlaygroundUseCase(repo repository.PlaygroundRepository,
	log utils.Logger) GetPlaygroundUseCase {

	return &getPlaygroundUseCaseImpl{repo: repo,
		log: log}
}

func (c *getPlaygroundUseCaseImpl) GetPlayground() *dto.PlaygroundDTO {
	playground, _ := c.repo.Get(false)
	defer c.repo.Save(playground, false)

	PlaygroundDTO := mappers.MapToPlaygroundDTO(playground)
	return PlaygroundDTO
}
