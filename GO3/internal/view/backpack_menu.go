package view

import (
	"s21_rogue/internal/app/usecase"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	view_utils "s21_rogue/internal/view/utils"

	"github.com/gdamore/tcell/v3"
)

type BackpackMenuFieldType int

const (
	BackpackFood BackpackMenuFieldType = iota
	BackpackScroll
	BackpackElixir
	BackpackWeapon
	BackpackNothing
)

type backpackMenu struct {
	BaseHandler
	selected      BackpackMenuFieldType
	selectedIndex int

	useCase usecase.UseBackpackUseCase

	screen *tcell.Screen
	config ViewerConfig
	styles ViewerStyles
	log    utils.Logger
}

func NewBackpackMenu(screen *tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	useCase usecase.UseBackpackUseCase,
	log utils.Logger) Handler {

	return &backpackMenu{screen: screen,
		config:   config,
		styles:   styles,
		log:      log,
		selected: BackpackNothing,
		useCase:  useCase}
}

func (m *backpackMenu) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if gameInfo.GameStateDTO == dto.MovingDTO {
		if m.selected == BackpackNothing {
			switch action {
			case value_objects.FoodAction:
				m.selected = BackpackFood
				m.selectedIndex = 0
			case value_objects.ScrollAction:
				m.selected = BackpackScroll
				m.selectedIndex = 0
			case value_objects.ElixirAction:
				m.selected = BackpackElixir
				m.selectedIndex = 0
			case value_objects.WeaponAction:
				m.selected = BackpackWeapon
				m.selectedIndex = 0
			default:
				return m.BaseHandler.Handle(action, gameInfo, playground)
			}
		}
		behaviours := make([]dto.ItemBehaviour, 0)
		menuName := ""
		switch m.selected {
		case BackpackFood:
			for _, e := range playground.BackpackDTO.FoodsDTO {
				behaviours = append(behaviours, e)
			}
			menuName = "Food Menu"
		case BackpackScroll:
			for _, e := range playground.BackpackDTO.ScrollsDTO {
				behaviours = append(behaviours, e)
			}
			menuName = "Scroll Menu"
		case BackpackElixir:
			for _, e := range playground.BackpackDTO.ElixirsDTO {
				behaviours = append(behaviours, e)
			}
			menuName = "Elixir Menu"
		case BackpackWeapon:
			for _, e := range playground.BackpackDTO.WeaponsDTO {
				behaviours = append(behaviours, e)
			}
			menuName = "Weapon Menu"
		}

		return m.process(behaviours,
			menuName,
			action,
			gameInfo,
			playground)

	}

	return m.BaseHandler.Handle(action, gameInfo, playground)
}

func (m *backpackMenu) process(l []dto.ItemBehaviour, menuName string, action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	m.updateSelectedIndex(action, len(l))

	if m.selectedIndex < len(l) {
		elem := l[m.selectedIndex]
		if action == value_objects.StartAction {
			m.useCase.UseItem(elem.GetUUID())
			m.selectedIndex = 0
		}
	}

	if action == value_objects.StartAction || action == value_objects.TerminateAction {
		m.selected = BackpackNothing
		m.selectedIndex = 0
		return m.BaseHandler.Handle(value_objects.Nothing, gameInfo, playground)
	}

	(*m.screen).Clear()
	view_utils.FillBackground(*m.screen, m.config.Theme.Background)
	m.printMenu(l, menuName)
	(*m.screen).Show()
	return true
}

func (m *backpackMenu) updateSelectedIndex(action value_objects.Action, count int) {
	switch action {
	case value_objects.UpAction:
		m.selectedIndex = max(0, m.selectedIndex-1)
	case value_objects.DownAction:
		m.selectedIndex = min(count, m.selectedIndex+1)
	}
}

func (m *backpackMenu) printMenu(l []dto.ItemBehaviour, menuName string) {
	h1 := m.styles.Ui.H1
	btn := m.styles.Ui.Text
	selected := m.styles.Ui.Value
	y := 3

	curText := menuName
	curX := (constants.MapWidth - len(curText)) / 2
	for i, ch := range curText {
		(*m.screen).SetContent(curX+i, y, ch, nil, h1)
	}

	y += 3
	i := 0

	for _, item := range l {
		curText = item.ToString()

		if item.IsSelected() {
			curText = curText + " < selected"
		}

		line := btn
		if i == m.selectedIndex {
			curText = "> " + curText
			line = selected
		} else {
			curText = "  " + curText
		}

		curX = 3
		for i, ch := range curText {
			(*m.screen).SetContent(curX+i, y, ch, nil, line)
		}
		y += 1
		i++
	}
}
