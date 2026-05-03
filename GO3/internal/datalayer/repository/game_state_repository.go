package repository

import (
	"errors"
	"s21_rogue/internal/domain/value_objects"
)

var (
	ErrGameNotFound = errors.New("Game not found in the repository")
)

type GameStateRepository interface {
	Save(gameState value_objects.GameState)
	Get() value_objects.GameState
}

type gameStateCacheRepository struct {
	GameState value_objects.GameState
}

func NewGameStateRepositoryCache(gameState value_objects.GameState) GameStateRepository {
	return &gameStateCacheRepository{GameState: gameState}
}

func (r *gameStateCacheRepository) Save(gameState value_objects.GameState) {
	r.GameState = gameState
}

func (r *gameStateCacheRepository) Get() value_objects.GameState {
	return r.GameState
}
