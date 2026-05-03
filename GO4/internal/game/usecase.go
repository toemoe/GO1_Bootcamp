package game

import (
	"errors"
	"tic-tac-toe/internal/domain/model"
	"tic-tac-toe/internal/domain/repository"
	"tic-tac-toe/internal/domain/service"

	"github.com/google/uuid"
)

type UseCase struct {
	repo    repository.GameRepository
	minimax service.GameService
}

func NewUseCase(repo repository.GameRepository, minimax service.GameService) *UseCase {
	return &UseCase{
		repo:    repo,
		minimax: minimax,
	}
}

func (u *UseCase) CreateGame() (model.Game, error) {
	game := model.Game{
		ID:    uuid.New(),
		Board: model.Board{},
	}
	if err := u.repo.SaveGame(game); err != nil {
		return model.Game{}, errors.New("failed to save game")
	}
	return game, nil
}

func (u *UseCase) MakeMove(gameId uuid.UUID, board model.Board) (model.Game, error) {
	var result model.Game
	err := u.repo.UpdateGame(gameId, func(game model.Game) (model.Game, error) {
		if u.minimax.IsGameOver(game.Board) {
			return model.Game{}, errors.New("game is already over")
		}

		if !u.minimax.ValidateBoard(game.Board, board) {
			return model.Game{}, errors.New("invalid board")
		}

		if u.minimax.IsGameOver(board) {
			game.Board = board
			result = game
			return game, nil
		}

		game.Board = board
		row, col := u.minimax.NextMove(game.Board)
		game.Board[row][col] = model.O
		result = game
		return game, nil
	})

	if err != nil {
		return model.Game{}, err
	}
	return result, nil
}
