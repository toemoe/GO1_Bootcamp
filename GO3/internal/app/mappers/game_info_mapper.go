package mappers

import (
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/value_objects"
)

func MapToGameInfoDTO(gameState value_objects.GameState, level int) *dto.GameInfoDTO {
	gameInfoDTO := dto.GameInfoDTO{}
	gameInfoDTO.Level = level
	switch gameState {
	case value_objects.Start:
		gameInfoDTO.GameStateDTO = dto.StartDTO
	case value_objects.NewGame:
		gameInfoDTO.GameStateDTO = dto.NewGameDTO
	case value_objects.LoadGame:
		gameInfoDTO.GameStateDTO = dto.LoadGameDTO
	case value_objects.Spawn:
		gameInfoDTO.GameStateDTO = dto.SpawnDTO
	case value_objects.Moving:
		gameInfoDTO.GameStateDTO = dto.MovingDTO
	case value_objects.Pause:
		gameInfoDTO.GameStateDTO = dto.PauseDTO
	case value_objects.GameOver:
		gameInfoDTO.GameStateDTO = dto.GameOverDTO
	case value_objects.GameWin:
		gameInfoDTO.GameStateDTO = dto.GameWinDTO
	case value_objects.ExitState:
		gameInfoDTO.GameStateDTO = dto.ExitStateDTO
	case value_objects.StatisticsState:
		gameInfoDTO.GameStateDTO = dto.StatisticsStateDTO
	}

	return &gameInfoDTO
}
