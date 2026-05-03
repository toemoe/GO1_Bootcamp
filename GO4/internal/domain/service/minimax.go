package service

import "tic-tac-toe/internal/domain/model"

type MinimaxService struct{}

func NewMinimaxService() *MinimaxService {
	return &MinimaxService{}
}

func (s *MinimaxService) NextMove(board model.Board) (int, int) {
	if s.IsGameOver(board) {
		return -1, -1
	}

	bestScore, bestRow, bestCol := 1001, -1, -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != model.Empty {
				continue
			}
			board[i][j] = model.O
			score := s.minimax(board, 0, true)
			board[i][j] = model.Empty
			if score < bestScore {
				bestScore, bestRow, bestCol = score, i, j
			}
		}
	}
	return bestRow, bestCol
}

func (s *MinimaxService) minimax(board model.Board, depth int, isMax bool) int {
	if s.checkWinner(board, model.X) {
		return 10 - depth
	} else if s.checkWinner(board, model.O) {
		return depth - 10
	} else if s.isBoardFull(board) {
		return 0
	}

	player, best := model.O, 1000
	if isMax {
		player, best = model.X, -1000
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != model.Empty {
				continue
			}
			board[i][j] = player
			score := s.minimax(board, depth+1, !isMax)
			board[i][j] = model.Empty

			if isMax && score > best || !isMax && score < best {
				best = score
			}
		}
	}
	return best
}

func (s *MinimaxService) ValidateBoard(old, updated model.Board) bool {
	diff := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if old[i][j] != updated[i][j] {
				if old[i][j] != model.Empty || updated[i][j] != model.X {
					return false
				}
				diff++
			}
		}
	}
	return diff == 1
}

func (s *MinimaxService) isBoardFull(board model.Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == model.Empty {
				return false
			}
		}
	}

	return true
}

func (s *MinimaxService) checkWinner(board model.Board, player int) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player ||
		board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

func (s *MinimaxService) IsGameOver(board model.Board) bool {
	return s.checkWinner(board, model.X) || s.checkWinner(board, model.O) || s.isBoardFull(board)
}
