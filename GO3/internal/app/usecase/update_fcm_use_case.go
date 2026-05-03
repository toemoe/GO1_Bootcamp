package usecase

import (
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/services"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
)

type UpdateFcmUseCase interface {
	Update(value_objects.Action) bool
}

type updateFcmUseCaseImpl struct {
	playgroundGeneratorService services.PlaygroundGeneratorService
	loadGameService            services.LoadGameService
	newGameService             services.NewGameService
	gameStateRepo              repository.GameStateRepository
	statisticsRepo             repository.StatisticsRepository
	playgroundRepo             repository.PlaygroundRepository
	publisher                  events.EventPublisher
	log                        utils.Logger
}

func NewUpdateFcmUseCase(
	playgroundGeneratorService services.PlaygroundGeneratorService,
	loadGameService services.LoadGameService,
	newGameService services.NewGameService,
	gameStateRepo repository.GameStateRepository,
	statisticsRepo repository.StatisticsRepository,
	playgroundRepo repository.PlaygroundRepository,
	publisher events.EventPublisher,
	log utils.Logger) UpdateFcmUseCase {

	return &updateFcmUseCaseImpl{
		playgroundGeneratorService: playgroundGeneratorService,
		loadGameService:            loadGameService,
		newGameService:             newGameService,
		gameStateRepo:              gameStateRepo,
		playgroundRepo:             playgroundRepo,
		statisticsRepo:             statisticsRepo,
		publisher:                  publisher,
		log:                        log}
}

func (c *updateFcmUseCaseImpl) Update(action value_objects.Action) bool {
	gameState := c.gameStateRepo.Get()

	switch gameState {
	case value_objects.Start:
		c.onStartState(action, &gameState)
	case value_objects.NewGame:
		c.onNewGameState(&gameState)
	case value_objects.LoadGame:
		c.onLoadGameState(&gameState)
	case value_objects.Spawn:
		c.onSpawnState(&gameState)
	case value_objects.Moving:
		c.onMovingState(action, &gameState)
	case value_objects.Pause:
		c.onPauseState(action, &gameState)
	case value_objects.StatisticsState:
		c.onStatisticsState(action, &gameState)
	case value_objects.GameOver:
		c.onGameOverState(action, &gameState)
	case value_objects.GameWin:
		c.onGameOverState(action, &gameState)
	}

	c.gameStateRepo.Save(gameState)

	return gameState != value_objects.ExitState
}

func (c *updateFcmUseCaseImpl) onStatisticsState(action value_objects.Action, gameState *value_objects.GameState) {
	if action == value_objects.BackAction {
		*gameState = value_objects.Start
	}
}

func (c *updateFcmUseCaseImpl) onStartState(action value_objects.Action, gameState *value_objects.GameState) {
	switch action {
	case value_objects.StartAction:
		*gameState = value_objects.NewGame
	case value_objects.LoadAction:
		*gameState = value_objects.LoadGame
	case value_objects.TerminateAction:
		*gameState = value_objects.ExitState
	case value_objects.StatisticsAction:
		*gameState = value_objects.StatisticsState
	}
}

func (c *updateFcmUseCaseImpl) onNewGameState(gameState *value_objects.GameState) {
	c.newGameService.NewGame()
	*gameState = value_objects.Moving
}

func (c *updateFcmUseCaseImpl) onLoadGameState(gameState *value_objects.GameState) {
	if c.loadGameService.LoadGame() {
		*gameState = value_objects.Moving
		return
	}
	*gameState = value_objects.NewGame
}

func (c *updateFcmUseCaseImpl) onSpawnState(gameState *value_objects.GameState) {
	playground, _ := c.playgroundRepo.Get(false)
	level := c.statisticsRepo.GetLevelById(playground.Id)

	newPlayground := c.playgroundGeneratorService.GenerateNewLevel(playground, level)
	c.playgroundRepo.Save(newPlayground, true)

	c.publisher.Notify(events.NewEvent(events.LevelChangedEvent))
	*gameState = value_objects.Moving
}

func (c *updateFcmUseCaseImpl) onMovingState(action value_objects.Action, gameState *value_objects.GameState) {
	ev := make([]events.Event, 0)
	if action == value_objects.TerminateAction {
		*gameState = value_objects.Pause
		return
	}
	if !action.IsMovingAction() {
		return
	}

	playground, _ := c.playgroundRepo.Get(false)
	nextCharacterPos := playground.Character.GetNextActionPos(action)

	if playground.Dungeon.InDungeon(nextCharacterPos) && playground.IsPortalPosition(nextCharacterPos) {
		*gameState = value_objects.Spawn
		level := c.statisticsRepo.GetLevelById(playground.Id)
		if level+1 > 21 {
			*gameState = value_objects.GameWin
		}
		c.publisher.Notify(events.NewEvent(events.NextLevelEvent))
		return
	}

	evs := playground.AttackOrUpdateCharacterPosition(nextCharacterPos)

	isTreasureFound := false
	for _, se := range evs {
		if se.GetType() == events.TreasureFoundEvent {
			isTreasureFound = true
			break
		}
	}
	if isTreasureFound {
		tev := make([]events.Event, 0)
		for range c.statisticsRepo.GetLevelById(playground.Id) {
			tev = append(tev, events.NewEvent(events.TreasureFoundEvent))
		}
		ev = append(ev, tev...)
	}

	ev = append(ev, evs...)

	evs2 := playground.AttackOrUpdateMonstersPosition()
	ev = append(ev, evs2...)

	if playground.Character.IsDeath() {
		*gameState = value_objects.GameOver
	}

	c.playgroundRepo.Save(playground, false)
	for _, e := range ev {
		c.publisher.Notify(e)
	}
}

func (c *updateFcmUseCaseImpl) onPauseState(action value_objects.Action, gameState *value_objects.GameState) {
	switch action {
	case value_objects.StartAction:
		*gameState = value_objects.Start
	case value_objects.TerminateAction:
		*gameState = value_objects.ExitState
	case value_objects.BackAction:
		*gameState = value_objects.Moving
	}
}

func (c *updateFcmUseCaseImpl) onGameOverState(action value_objects.Action, gameState *value_objects.GameState) {
	switch action {
	case value_objects.StartAction:
		*gameState = value_objects.Start
	case value_objects.TerminateAction:
		*gameState = value_objects.ExitState
	}
}
