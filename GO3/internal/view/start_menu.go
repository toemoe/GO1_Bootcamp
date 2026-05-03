package view

import (
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	view_utils "s21_rogue/internal/view/utils"

	"github.com/gdamore/tcell/v3"
)

type MenuFieldType int

const (
	NewGameMenuFieldType    MenuFieldType = iota
	LoadGameMenuFieldType                 = iota
	StatisticsMenuFieldType               = iota
	ExitMenuFieldType
)

type MenuField struct {
	Type MenuFieldType
	Text string
}

type menu struct {
	BaseHandler
	selected int
	items    []MenuField

	screen *tcell.Screen
	config ViewerConfig
	styles ViewerStyles
	log    utils.Logger
}

func NewStartMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	log utils.Logger) Handler {
	items := make([]MenuField, 0)
	items = append(items, MenuField{Type: NewGameMenuFieldType, Text: "New Game"})
	items = append(items, MenuField{Type: LoadGameMenuFieldType, Text: "Load Game"})
	items = append(items, MenuField{Type: StatisticsMenuFieldType, Text: "Statistics"})
	items = append(items, MenuField{Type: ExitMenuFieldType, Text: "Exit"})

	return &menu{screen: screen,
		config:   config,
		styles:   styles,
		log:      log,
		items:    items,
		selected: 0}
}

func (m *menu) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if gameInfo.GameStateDTO == dto.StartDTO {
		switch action {
		case value_objects.UpAction:
			m.selected = max(0, m.selected-1)
		case value_objects.DownAction:
			m.selected = min(len(m.items)-1, m.selected+1)
		case value_objects.StartAction:
			switch m.items[m.selected].Type {
			case NewGameMenuFieldType:
				m.selected = 0
				return m.BaseHandler.Handle(value_objects.StartAction, gameInfo, playground)
			case LoadGameMenuFieldType:
				m.selected = 0
				return m.BaseHandler.Handle(value_objects.LoadAction, gameInfo, playground)
			case StatisticsMenuFieldType:
				return m.BaseHandler.Handle(value_objects.StatisticsAction, gameInfo, playground)
			case ExitMenuFieldType:
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

func (m *menu) printMainMenu() {
	h1 := m.styles.Ui.H1
	btn := m.styles.Ui.Text
	y := constants.MapHeight / 6

	curText := "ROGUE"
	curX := (constants.MapWidth - len(curText)) / 2
	for i, ch := range curText {
		(*m.screen).SetContent(curX+i, y, ch, nil, h1)
	}

	y += 4
	for i := range m.items {
		curText = m.items[i].Text
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
