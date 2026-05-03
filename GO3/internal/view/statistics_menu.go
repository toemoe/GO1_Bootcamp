package view

import (
	"fmt"
	"s21_rogue/internal/app/usecase"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	view_utils "s21_rogue/internal/view/utils"

	"github.com/gdamore/tcell/v3"
)

const (
	MaxLenStatisticsField = 20
	MaxLenStatisticsValue = 5
)

type statisticsMenu struct {
	BaseHandler
	cache    []*dto.StatisticsDTO
	isInit   bool
	selected int
	opened   bool

	useCase usecase.GetStatisticsUseCase

	screen *tcell.Screen
	config ViewerConfig
	styles ViewerStyles
	log    utils.Logger
}

func NewStatisticsMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	log utils.Logger,
	useCase usecase.GetStatisticsUseCase) Handler {
	return &statisticsMenu{
		screen:  screen,
		config:  config,
		styles:  styles,
		log:     log,
		useCase: useCase,

		selected: 0,
		opened:   false,
		isInit:   false}
}

func (m *statisticsMenu) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if gameInfo.GameStateDTO == dto.StatisticsStateDTO {
		if !m.isInit {
			m.cache = m.useCase.GetStatistics()
			m.selected = 0
			m.opened = false
			m.isInit = true
		}

		switch action {
		case value_objects.UpAction:
			m.selected = max(0, m.selected-1)
		case value_objects.DownAction:
			m.selected = min(len(m.cache)-1, m.selected+1)
		case value_objects.StartAction:
			m.opened = true
		case value_objects.TerminateAction:
			if m.opened {
				m.opened = false
			} else {
				m.isInit = false
				return m.BaseHandler.Handle(value_objects.BackAction, gameInfo, playground)
			}
		}
	} else {
		return m.BaseHandler.Handle(action, gameInfo, playground)
	}

	(*m.screen).Clear()
	view_utils.FillBackground(*m.screen, m.config.Theme.Background)
	m.printMenu()
	(*m.screen).Show()
	return true
}

func (m *statisticsMenu) printMenu() {
	h1 := m.styles.Ui.H1
	btn := m.styles.Ui.Text
	y := 3

	curText := "LeaderBoard"
	curX := (constants.MapWidth - len(curText)) / 2
	for i, ch := range curText {
		(*m.screen).SetContent(curX+i, y, ch, nil, h1)
	}

	y += 3
	curX = 3

	if !m.opened {
		for i, s := range m.cache {
			chooseCh := ">"
			if i != m.selected {
				chooseCh = " "
			}
			curText = fmt.Sprintf("%v %v\t%v", chooseCh, s.Id, s.TreasureFound)

			for i, ch := range curText {
				(*m.screen).SetContent(curX+i, y, ch, nil, btn)
			}
			y += 1
		}
	} else {
		s := m.cache[m.selected]
		var fields = []struct {
			text  string
			value int
		}{
			{"Level", s.Level},
			{"Monster dead", s.MonsterDead},
			{"Given damage", s.GivenDamage},
			{"Recieved damage", s.RecievedDamage},
			{"Food eaten", s.FoodEaten},
			{"Elixir drunk", s.ElixirsDrunk},
			{"Scrolls read", s.ScrollsRead},
			{"Treasure found", s.TreasureFound},
			{"Stepped count", s.SteppedCount},
		}

		for _, v := range fields {
			curText = fmt.Sprintf("%-*s - %+*v", MaxLenStatisticsField, v.text, MaxLenStatisticsValue, v.value)
			for i, ch := range curText {
				(*m.screen).SetContent(curX+i, y, ch, nil, btn)
			}
			y += 1
		}

	}

}
