package api

import (
	domain "tic-tac-toe/internal/domain/model"
)

func ToDTO(game domain.Game) GameDTO {
	return GameDTO{
		ID:    game.ID,
		Board: game.Board,
	}
}

func ToDomain(dto GameDTO) domain.Game {
	return domain.Game{
		ID:    dto.ID,
		Board: dto.Board,
	}
}
