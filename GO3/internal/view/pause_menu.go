package view

import (
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	view_utils "s21_rogue/internal/view/utils"

	"github.com/gdamore/tcell/v3"
)

type childMenu struct {
	BaseHandler
	selected int
	items    []string
	mainDTO  string
	label    string

	screen *tcell.Screen
	config ViewerConfig
	styles ViewerStyles
	log    utils.Logger
}

const (
	continueTextMenu string = "Continue"
	newGameTextMenu  string = "New Game"
	exitTextMenu     string = "Exit"
)

func NewPauseMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	log utils.Logger) Handler {
	items := make([]string, 0)
	items = append(items, continueTextMenu)
	items = append(items, newGameTextMenu)
	items = append(items, exitTextMenu)

	return &childMenu{screen: screen,
		config: config,
		styles: styles,
		log:    log,
		items:  items,

		selected: 0,
		mainDTO:  dto.PauseDTO,
		label:    "Pause"}
}

func NewGameOverMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	log utils.Logger) Handler {
	items := make([]string, 0)
	items = append(items, newGameTextMenu)
	items = append(items, exitTextMenu)

	return &childMenu{screen: screen,
		config: config,
		styles: styles,
		log:    log,
		items:  items,

		selected: 0,
		mainDTO:  dto.GameOverDTO,
		label:    "Game Over"}
}
func NewGameWinMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	log utils.Logger) Handler {
	items := make([]string, 0)
	items = append(items, newGameTextMenu)
	items = append(items, exitTextMenu)

	return &childMenu{screen: screen,
		config: config,
		styles: styles,
		log:    log,
		items:  items,

		selected: 0,
		mainDTO:  dto.GameWinDTO,
		label:    "Game Win"}
}

func (m *childMenu) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if gameInfo.GameStateDTO == m.mainDTO {
		switch action {
		case value_objects.UpAction:
			m.selected = max(0, m.selected-1)
		case value_objects.DownAction:
			m.selected = min(len(m.items)-1, m.selected+1)
		case value_objects.StartAction:
			switch m.items[m.selected] {
			case continueTextMenu:
				m.selected = 0
				return m.BaseHandler.Handle(value_objects.BackAction, gameInfo, playground)
			case newGameTextMenu:
				m.selected = 0
				return m.BaseHandler.Handle(value_objects.StartAction, gameInfo, playground)
			case exitTextMenu:
				return m.BaseHandler.Handle(value_objects.TerminateAction, gameInfo, playground)
			}
		}
	} else {
		return m.BaseHandler.Handle(action, gameInfo, playground)
	}

	(*m.screen).Clear()
	view_utils.FillBackground(*m.screen, m.config.Theme.Background)
	m.printMainMenu()
	(*m.screen).Show()
	return true
}

func (m *childMenu) printMainMenu() {
	h1 := m.styles.Ui.H1
	btn := m.styles.Ui.Text
	y := constants.MapHeight / 6

	curText := m.label
	curX := (constants.MapWidth - len(curText)) / 2
	for i, ch := range curText {
		(*m.screen).SetContent(curX+i, y, ch, nil, h1)
	}

	y += 4
	for i := range m.items {
		curText = m.items[i]
		if i == m.selected {
			curText = "-- " + curText + " --"
		}
		curX = (constants.MapWidth - len(curText)) / 2
		for i, ch := range curText {
			(*m.screen).SetContent(curX+i, y, ch, nil, btn)
		}
		y += 2
	}
}
