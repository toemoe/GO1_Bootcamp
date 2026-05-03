package main

import (
	"os"
	"path/filepath"
	"s21_rogue/internal/app/usecase"
	"s21_rogue/internal/datalayer/repository"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/services"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/utils/generators"
	"s21_rogue/internal/domain/utils/validators"
	"s21_rogue/internal/domain/value_objects"
	"s21_rogue/internal/view"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func main() {
	// Init logger
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Join(homeDir, "s21_rogue")

	err = os.MkdirAll(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	filePath := filepath.Join(dirPath, "info.log")

	logger, err := utils.NewFileLogger(filePath)
	if err != nil {
		panic(err)
	}

	config := view.BasicViewerConfig()
	styles := view.BasicViewerStyles()

	defer logger.Sync()

	// InitDomain
	playgroundRepo := repository.NewPlaygroundRepositoryJson(filepath.Join(dirPath, "playground.json"), logger)
	statisticsRepo := repository.NewStatisticsJsonRepository(filepath.Join(dirPath, "statistics.json"), logger)
	gameStateRepo := repository.NewGameStateRepositoryCache(value_objects.Start)

	publisher := events.NewEventPublisher()

	validate := validator.New()
	validate.RegisterStructValidation(validators.RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(validators.DungeonStructLevelValidation, entities.Dungeon{})

	dg := generators.NewDungeonGenerator(logger)
	playgroundService := services.NewPlaygroundGeneratorService(dg, logger)
	newGameService := services.NewNewGameService(playgroundService, playgroundRepo, statisticsRepo, publisher, logger)
	loadGameService := services.NewLoadGameService(playgroundRepo, statisticsRepo, publisher, validate, logger)

	newGameService.NewGame()

	updateFcmUseCase := usecase.NewUpdateFcmUseCase(playgroundService, loadGameService, newGameService, gameStateRepo, statisticsRepo, playgroundRepo, publisher, logger)
	useBackpackItemUseCase := usecase.NewUseBackpackUseCase(playgroundRepo, publisher, logger)
	getGameInfoUseCase := usecase.NewGetGameInfoUseCase(gameStateRepo, statisticsRepo, logger)
	getPlaygroundUseCase := usecase.NewGetPlaygroundUseCase(playgroundRepo, logger)
	getStatisticsUseCase := usecase.NewGetStatisticsUseCase(statisticsRepo, logger)

	// init Viewer
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	screen.Clear()

	defer screen.Fini()

	viewer := view.NewTcellViewer(screen, config, styles, updateFcmUseCase, logger)
	statisticsMenu := view.NewStatisticsMenu(&screen, config, styles, logger, getStatisticsUseCase)
	menu := view.NewStartMenu(&screen, config, styles, logger)
	pauseMenu := view.NewPauseMenu(&screen, config, styles, logger)
	backpackMenu := view.NewBackpackMenu(&screen, config, styles, useBackpackItemUseCase, logger)
	gameOverMenu := view.NewGameOverMenu(&screen, config, styles, logger)
	gameWinMenu := view.NewGameWinMenu(&screen, config, styles, logger)

	keyboard := view.NewKeyboard(&screen)

	menu.Next(statisticsMenu)
	statisticsMenu.Next(pauseMenu)
	pauseMenu.Next(gameOverMenu)
	gameOverMenu.Next(gameWinMenu)
	gameWinMenu.Next(backpackMenu)
	backpackMenu.Next(viewer)

	actionChan := keyboard.GetActionChan()
	isTerminate := false
	for !isTerminate {
		var sig value_objects.Action
		select {
		case temp := <-actionChan:
			sig = temp
		case <-time.Tick(50 * time.Millisecond):
			sig = value_objects.Nothing
		}

		playgroundDTO := getPlaygroundUseCase.GetPlayground()

		id, err := uuid.Parse(playgroundDTO.UUID)
		if err != nil {
			panic(err)
		}
		gameInfoDTO := getGameInfoUseCase.GetGameInfo(id)
		isTerminate = !menu.Handle(sig, gameInfoDTO, playgroundDTO)
	}

	logger.Info("Terminated")
}
