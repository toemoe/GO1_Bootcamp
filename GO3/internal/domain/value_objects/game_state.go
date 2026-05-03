package value_objects

type GameState int

const (
	Start GameState = iota
	NewGame
	LoadGame
	Spawn
	Moving
	Pause
	GameOver
	GameWin
	ExitState
	StatisticsState
)
