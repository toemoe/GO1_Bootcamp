package repository

import (
	"tic-tac-toe/internal/domain/model"

	"github.com/google/uuid"
)

type GameRepository interface {
	SaveGame(game model.Game) error
	GetGame(id uuid.UUID) (model.Game, error)
	UpdateGame(id uuid.UUID, update func(game model.Game) (model.Game, error)) error
}
