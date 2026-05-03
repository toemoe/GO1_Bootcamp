package dto

const (
	StartDTO           string = "start"
	NewGameDTO         string = "new_game"
	LoadGameDTO        string = "load_game"
	SpawnDTO           string = "spawn"
	MovingDTO          string = "moving"
	PauseDTO           string = "pause"
	GameOverDTO        string = "game_over"
	GameWinDTO         string = "game_win"
	ExitStateDTO       string = "exit"
	StatisticsStateDTO string = "statistics"
)

type GameInfoDTO struct {
	GameStateDTO string `json:"game_state"`
	Level        int    `json:"level"`
}
