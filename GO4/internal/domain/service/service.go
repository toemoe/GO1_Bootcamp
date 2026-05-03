package service

import "tic-tac-toe/internal/domain/model"

type GameService interface {
	NextMove(board model.Board) (int, int)
	ValidateBoard(old, new model.Board) bool
	IsGameOver(board model.Board) bool
}
