package storage

import (
	domain "tic-tac-toe/internal/domain/model"
)

func ToEntity(game domain.Game) GameEntity {
	return GameEntity{
		ID:    game.ID,
		Board: game.Board,
	}
}

func ToDomain(e GameEntity) domain.Game {
	return domain.Game{
		ID:    e.ID,
		Board: e.Board,
	}
}
