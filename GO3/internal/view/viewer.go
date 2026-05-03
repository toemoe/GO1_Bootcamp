package view

import (
	"fmt"
	"math"
	"s21_rogue/internal/app/usecase"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	domain_utils "s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	view_utils "s21_rogue/internal/view/utils"
	"strconv"

	"github.com/gdamore/tcell/v3"
)

const (
	OffsetGameplayOutline = 2
	OffsetWallOutline     = 2
	OffsetCommonGameplay  = OffsetGameplayOutline + OffsetWallOutline
)

type tcellViewer struct {
	BaseHandler
	useCase    usecase.UpdateFcmUseCase
	gameInfo   *dto.GameInfoDTO
	playground *dto.PlaygroundDTO
	log        domain_utils.Logger
	screen     tcell.Screen
	config     ViewerConfig
	styles     ViewerStyles
}

func NewTcellViewer(screen tcell.Screen,
	config ViewerConfig,
	styles ViewerStyles,
	useCase usecase.UpdateFcmUseCase,
	log domain_utils.Logger) Handler {

	log.Info("Initial viewer")
	return &tcellViewer{log: log,
		screen:   screen,
		gameInfo: nil,
		config:   config,
		styles:   styles,
		useCase:  useCase}
}

func (v *tcellViewer) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if gameInfo.GameStateDTO == dto.MovingDTO {
		v.screen.Clear()
		v.gameInfo = gameInfo
		v.playground = playground
		view_utils.FillBackground(v.screen, v.config.Theme.Background)

		view_utils.PrintRectangle(v.screen, &v.styles.MapStyles.Current, 0, constants.MapHeight+OffsetCommonGameplay, 0, constants.MapWidth+OffsetCommonGameplay)

		v.printCorridors()
		v.printRooms()
		v.printPortal()
		v.printMonsters()
		v.printFood()
		v.printScroll()
		v.printElixir()
		v.printWeapon()
		v.printTreasure()
		v.printCharacter()

		v.printMessagePanel()
		v.printToolBar()
		v.printHUD()
		v.screen.Show()
	}

	return v.useCase.Update(action)

}

func (v *tcellViewer) inShowedRadius(pos value_objects.Position) bool {
	return math.Sqrt(float64((pos.X-v.playground.CharacterDTO.Position.X)*(pos.X-v.playground.CharacterDTO.Position.X)+(pos.Y-v.playground.CharacterDTO.Position.Y)*(pos.Y-v.playground.CharacterDTO.Position.Y))) < constants.RadiusFogOfWar
}

func (v *tcellViewer) offsetToPlaygroundViewer(pos value_objects.Position) *value_objects.Position {
	return value_objects.NewPosition(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline)

}

func (v *tcellViewer) printRooms() {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			room := v.playground.DungeonDTO.Rooms[i][j]

			switch room.VisitStateDTO {
			case "current_space":
				v.printRoom(room, v.styles.MapStyles.Current)
			case "visited":
				v.printRoom(room, v.styles.MapStyles.Visited)
			case "not_visited":
				v.printNotVisitedRoom(room)
			}

		}
	}
}

func (v *tcellViewer) printRoom(room dto.RoomDTO, style tcell.Style) {
	view_utils.PrintRectangle(v.screen, &style, constants.MapHeight-(room.TopLeft.Y+1)+OffsetWallOutline,
		constants.MapHeight-(room.BotRight.Y-1)+OffsetWallOutline,
		(room.TopLeft.X-1)+OffsetWallOutline, (room.BotRight.X+1)+OffsetWallOutline)

	for _, coord := range room.Doors {
		if !(coord.X == value_objects.UninitializedCoord && coord.Y == value_objects.UninitializedCoord) {
			v.screen.SetContent(coord.X+OffsetWallOutline, constants.MapHeight-coord.Y+OffsetWallOutline, '#', nil, style)
		}
	}
}

func (v *tcellViewer) printNotVisitedRoom(room dto.RoomDTO) {
	style := v.styles.MapStyles.Visited

	for x := room.TopLeft.X - 1; x < room.BotRight.X+1; x++ {
		checkPos := value_objects.NewPosition(x, room.TopLeft.Y+1)
		if v.inShowedRadius(*checkPos) {
			pgCoord := v.offsetToPlaygroundViewer(*checkPos)
			v.screen.SetContent(pgCoord.X, pgCoord.Y, '-', nil, style)
		}

		checkPos = value_objects.NewPosition(x, room.BotRight.Y-1)
		if v.inShowedRadius(*checkPos) {
			pgCoord := v.offsetToPlaygroundViewer(*checkPos)
			v.screen.SetContent(pgCoord.X, pgCoord.Y, '-', nil, style)
		}
	}

	for y := room.BotRight.Y - 1; y < room.TopLeft.Y+1; y++ {
		checkPos := value_objects.NewPosition(room.TopLeft.X-1, y)
		if v.inShowedRadius(*checkPos) {
			pgCoord := v.offsetToPlaygroundViewer(*checkPos)
			v.screen.SetContent(pgCoord.X, pgCoord.Y, '|', nil, style)
		}

		checkPos = value_objects.NewPosition(room.BotRight.X+1, y)
		if v.inShowedRadius(*checkPos) {
			pgCoord := v.offsetToPlaygroundViewer(*checkPos)
			v.screen.SetContent(pgCoord.X, pgCoord.Y, '|', nil, style)
		}
	}

	checkPos := value_objects.NewPosition(room.TopLeft.X-1, room.TopLeft.Y+1)
	if v.inShowedRadius(*checkPos) {
		pgCoord := v.offsetToPlaygroundViewer(*checkPos)
		v.screen.SetContent(pgCoord.X, pgCoord.Y, '┌', nil, style)
	}

	checkPos = value_objects.NewPosition(room.BotRight.X+1, room.TopLeft.Y+1)
	if v.inShowedRadius(*checkPos) {
		pgCoord := v.offsetToPlaygroundViewer(*checkPos)
		v.screen.SetContent(pgCoord.X, pgCoord.Y, '┐', nil, style)
	}

	checkPos = value_objects.NewPosition(room.TopLeft.X-1, room.BotRight.Y-1)
	if v.inShowedRadius(*checkPos) {
		pgCoord := v.offsetToPlaygroundViewer(*checkPos)
		v.screen.SetContent(pgCoord.X, pgCoord.Y, '└', nil, style)
	}

	checkPos = value_objects.NewPosition(room.BotRight.X+1, room.BotRight.Y-1)
	if v.inShowedRadius(*checkPos) {
		pgCoord := v.offsetToPlaygroundViewer(*checkPos)
		v.screen.SetContent(pgCoord.X, pgCoord.Y, '┘', nil, style)
	}

	for _, coord := range room.Doors {
		if !(coord.X == value_objects.UninitializedCoord && coord.Y == value_objects.UninitializedCoord) {
			if v.inShowedRadius(coord) {
				pgCoord := v.offsetToPlaygroundViewer(coord)
				v.screen.SetContent(pgCoord.X, pgCoord.Y, '#', nil, style)
			}
		}
	}
}

func (v *tcellViewer) printCorridors() {
	for _, corridor := range v.playground.DungeonDTO.CorridorsDTO {
		switch corridor.VisitStateDTO {
		case "current_space":
			v.printCorridor(&corridor, v.styles.MapStyles.Current)
		case "visited":
			v.printCorridor(&corridor, v.styles.MapStyles.Visited)
		case "not_visited":
			v.printCorridor(&corridor, v.styles.MapStyles.Visited)
			v.blurUnshowedCorridorWall(&corridor)
		}
	}
}

func (v *tcellViewer) blurUnshowedCorridorWall(corridor *dto.CorridorDTO) {
	coords := view_utils.GetDilateCorridorCoords(corridor)
	for _, c := range coords {
		if !v.inShowedRadius(*c) {
			pgCoord := v.offsetToPlaygroundViewer(*c)
			v.screen.SetContent(pgCoord.X, pgCoord.Y, ' ', nil, v.styles.MapStyles.NotVisited)
		}
	}
}

func (v *tcellViewer) printCorridor(corridor *dto.CorridorDTO, style tcell.Style) {
	isDirectCorridor := view_utils.IsDirectCorridor(corridor)
	if isDirectCorridor && view_utils.IsHorizontalCorridor(corridor, isDirectCorridor) {
		for i := corridor.Positions[0].X; i <= corridor.Positions[3].X; i++ {
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y+1+OffsetWallOutline, '-', nil, style)
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y-1+OffsetWallOutline, '-', nil, style)
		}
	} else if !isDirectCorridor && view_utils.IsHorizontalCorridor(corridor, isDirectCorridor) {
		for i := corridor.Positions[0].X; i < corridor.Positions[1].X-1; i++ {
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y+1+OffsetWallOutline, '-', nil, style)
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y-1+OffsetWallOutline, '-', nil, style)
		}
		for i := corridor.Positions[2].X + 2; i <= corridor.Positions[3].X; i++ {
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[3].Y+1+OffsetWallOutline, '-', nil, style)
			v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[3].Y-1+OffsetWallOutline, '-', nil, style)
		}

		if corridor.Positions[1].Y > corridor.Positions[2].Y {
			v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '┐', nil, style)
			v.screen.SetContent(corridor.Positions[2].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y+1+OffsetWallOutline, '└', nil, style)

			v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '┐', nil, style)
			v.screen.SetContent(corridor.Positions[2].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y-1+OffsetWallOutline, '└', nil, style)

			for i := corridor.Positions[0].X; i <= corridor.Positions[1].X; i++ {
				v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y-1+OffsetWallOutline, '-', nil, style)
			}

			for i := corridor.Positions[2].X; i < corridor.Positions[3].X; i++ {
				v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[3].Y+1+OffsetWallOutline, '-', nil, style)
			}

			for i := corridor.Positions[2].Y + 1; i < corridor.Positions[1].Y; i++ {
				v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-i+1+OffsetWallOutline, '|', nil, style)
				v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-i-1+OffsetWallOutline, '|', nil, style)
			}

		} else {
			v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '┘', nil, style)
			v.screen.SetContent(corridor.Positions[2].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y-1+OffsetWallOutline, '┌', nil, style)

			v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '┘', nil, style)
			v.screen.SetContent(corridor.Positions[2].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y+1+OffsetWallOutline, '┌', nil, style)

			for i := corridor.Positions[0].X; i <= corridor.Positions[1].X; i++ {
				v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[0].Y+1+OffsetWallOutline, '-', nil, style)
			}

			for i := corridor.Positions[2].X; i < corridor.Positions[3].X; i++ {
				v.screen.SetContent(i+OffsetWallOutline, constants.MapHeight-corridor.Positions[3].Y-1+OffsetWallOutline, '-', nil, style)
			}

			for i := corridor.Positions[1].Y + 1; i < corridor.Positions[2].Y; i++ {
				v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-i-1+OffsetWallOutline, '|', nil, style)
				v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-i+1+OffsetWallOutline, '|', nil, style)
			}
		}
	} else if isDirectCorridor && !view_utils.IsHorizontalCorridor(corridor, isDirectCorridor) {
		for i := corridor.Positions[0].Y; i <= corridor.Positions[3].Y; i++ {
			v.screen.SetContent(corridor.Positions[0].X-1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			v.screen.SetContent(corridor.Positions[0].X+1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
		}
	} else if !isDirectCorridor && !view_utils.IsHorizontalCorridor(corridor, isDirectCorridor) {
		for i := corridor.Positions[0].Y; i < corridor.Positions[1].Y-1; i++ {
			v.screen.SetContent(corridor.Positions[0].X+1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			v.screen.SetContent(corridor.Positions[0].X-1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
		}
		for i := corridor.Positions[2].Y + 2; i <= corridor.Positions[3].Y; i++ {
			v.screen.SetContent(corridor.Positions[3].X+1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			v.screen.SetContent(corridor.Positions[3].X-1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
		}

		if corridor.Positions[1].X > corridor.Positions[2].X {
			v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '┐', nil, style)
			v.screen.SetContent(corridor.Positions[2].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y+1+OffsetWallOutline, '└', nil, style)

			v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '┐', nil, style)
			v.screen.SetContent(corridor.Positions[2].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y-1+OffsetWallOutline, '└', nil, style)

			for i := corridor.Positions[0].Y; i <= corridor.Positions[1].Y; i++ {
				v.screen.SetContent(corridor.Positions[0].X+1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			}

			for i := corridor.Positions[2].Y; i < corridor.Positions[3].Y; i++ {
				v.screen.SetContent(corridor.Positions[3].X-1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			}

			for i := corridor.Positions[2].X + 1; i < corridor.Positions[1].X; i++ {
				v.screen.SetContent(i-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '-', nil, style)
				v.screen.SetContent(i+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '-', nil, style)
			}
		} else {
			v.screen.SetContent(corridor.Positions[1].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '┌', nil, style)
			v.screen.SetContent(corridor.Positions[2].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y-1+OffsetWallOutline, '┘', nil, style)

			v.screen.SetContent(corridor.Positions[1].X-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '┌', nil, style)
			v.screen.SetContent(corridor.Positions[2].X+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[2].Y+1+OffsetWallOutline, '┘', nil, style)

			for i := corridor.Positions[0].Y; i <= corridor.Positions[1].Y; i++ {
				v.screen.SetContent(corridor.Positions[0].X-1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			}

			for i := corridor.Positions[2].Y; i < corridor.Positions[3].Y; i++ {
				v.screen.SetContent(corridor.Positions[3].X+1+OffsetWallOutline, constants.MapHeight-i+OffsetWallOutline, '|', nil, style)
			}

			for i := corridor.Positions[1].X + 1; i < corridor.Positions[2].X; i++ {
				v.screen.SetContent(i+1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y+1+OffsetWallOutline, '-', nil, style)
				v.screen.SetContent(i-1+OffsetWallOutline, constants.MapHeight-corridor.Positions[1].Y-1+OffsetWallOutline, '-', nil, style)
			}

		}

	}
}

func (v *tcellViewer) printCharacter() {
	v.screen.SetContent(v.playground.CharacterDTO.Position.X+OffsetWallOutline,
		constants.MapHeight-v.playground.CharacterDTO.Position.Y+OffsetWallOutline, '@', nil, v.styles.Characters.Player)
}

func (v *tcellViewer) printPortal() {
	portalPos := value_objects.Position{
		X: v.playground.DungeonDTO.Portal.X,
		Y: v.playground.DungeonDTO.Portal.Y,
	}
	curr := view_utils.GetCurrentRoom(v.playground)
	if view_utils.IsInRoom(portalPos, curr) || v.inShowedRadius(portalPos) {
		v.screen.SetContent(portalPos.X+OffsetWallOutline,
			constants.MapHeight-portalPos.Y+OffsetWallOutline, '#', nil, v.styles.MapStyles.Portal)
	}
}

func (v *tcellViewer) printMonsters() {
	curr := view_utils.GetCurrentRoom(v.playground)
	type monsterRender struct {
		char  rune
		style tcell.Style
	}

	renderMap := map[dto.MonsterTypeDTO]monsterRender{
		dto.ZombieTypeDTO:  {'z', v.styles.Characters.Zombie},
		dto.GhostTypeDTO:   {'g', v.styles.Characters.Ghost},
		dto.VampireTypeDTO: {'v', v.styles.Characters.Vampire},
		dto.OgreTypeDTO:    {'o', v.styles.Characters.Ogre},
		dto.SnakeTypeDTO:   {'s', v.styles.Characters.Snake},
	}

	for _, monster := range v.playground.MonstersDTO {
		render := renderMap[dto.MonsterTypeDTO(monster.MonsterTypeDTO)]
		pos := monster.Position

		inRoom := curr != nil && view_utils.IsInRoom(pos, curr)
		inCorridor := false
		if !inRoom {
			for i := range v.playground.DungeonDTO.CorridorsDTO {
				if view_utils.IsInCorridor(pos, &v.playground.DungeonDTO.CorridorsDTO[i]) {
					inCorridor = true
					break
				}
			}
		}
		if inRoom || inCorridor || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, render.char, nil, render.style)
		}
	}
}

func (v *tcellViewer) printFood() {
	foods := v.playground.FoodsDTO

	curr := view_utils.GetCurrentRoom(v.playground)
	for _, food := range foods {
		pos := food.Position
		if view_utils.IsInRoom(pos, curr) || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, 'f', nil, v.styles.Items.Food)
		}
	}
}

func (v *tcellViewer) printScroll() {
	scrolls := v.playground.ScrollsDTO
	curr := view_utils.GetCurrentRoom(v.playground)
	for _, scroll := range scrolls {
		pos := scroll.Position
		if view_utils.IsInRoom(pos, curr) || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, 's', nil, v.styles.Items.Scroll)
		}
	}
}

func (v *tcellViewer) printElixir() {
	curr := view_utils.GetCurrentRoom(v.playground)
	elixis := v.playground.ElixirsDTO
	for _, elixir := range elixis {
		pos := elixir.Position
		if view_utils.IsInRoom(pos, curr) || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, 'e', nil, v.styles.Items.Elixir)
		}
	}
}

func (v *tcellViewer) printWeapon() {
	curr := view_utils.GetCurrentRoom(v.playground)
	weapons := v.playground.WeaponsDTO
	for _, weapon := range weapons {
		pos := weapon.Position
		if view_utils.IsInRoom(pos, curr) || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, 'w', nil, v.styles.Items.Weapon)
		}
	}
}

func (v *tcellViewer) printTreasure() {
	curr := view_utils.GetCurrentRoom(v.playground)
	treasures := v.playground.TreasuresDTO
	for _, treasure := range treasures {
		pos := treasure.Position
		if view_utils.IsInRoom(pos, curr) || v.inShowedRadius(pos) {
			v.screen.SetContent(pos.X+OffsetWallOutline, constants.MapHeight-pos.Y+OffsetWallOutline, 't', nil, v.styles.Items.Treasure)
		}
	}
}

func (v *tcellViewer) printMessagePanel() {
	x1, y1 := constants.MapWidth+OffsetCommonGameplay+5, 15
	x2, y2 := constants.MapWidth+OffsetCommonGameplay+40, 25
	view_utils.PrintRectangle(v.screen, &v.styles.MapStyles.Current, y1, y2, x1, x2)

	x := constants.MapWidth + OffsetCommonGameplay + 8
	y := 16

	for _, msg := range v.log.GetMessages() {
		shortMsg := fmt.Sprintf("%.*s", x2-x-2, msg)
		view_utils.PrintText(v.screen, x, y, shortMsg, v.styles.Ui.Warning)
		y += 2
	}
}

func (v *tcellViewer) printToolBar() {
	type toolBar struct {
		name  string
		value string
	}

	items := []toolBar{
		{"w", "- move up"},
		{"a", "- move left"},
		{"s", "- move down"},
		{"d", "- move right"},
		{"j", "- food menu"},
		{"e", "- scroll menu"},
		{"k", "- elixir menu"},
		{"h", "- weapon menu"},
		{"esc", "- exit"},
	}
	x1, y1 := constants.MapWidth+OffsetCommonGameplay+5, 1
	x2, y2 := constants.MapWidth+OffsetCommonGameplay+32, 13
	view_utils.PrintRectangle(v.screen, &v.styles.MapStyles.Current, y1, y2, x1, x2)

	x := constants.MapWidth + OffsetCommonGameplay + 11
	y := 2
	view_utils.PrintText(v.screen, x, y, "Game managemement", v.styles.Ui.Text)
	x += 1
	y += 2

	for _, item := range items {
		view_utils.PrintText(v.screen, x, y, item.name, v.styles.Ui.Value)
		view_utils.PrintText(v.screen, x+len(item.name)+1, y, item.value, v.styles.Ui.Text)
		y += 1
	}

}

func (v *tcellViewer) printHUD() {
	type hud struct {
		name  string
		value string
	}

	items := []hud{
		{"Level: ", strconv.Itoa(v.gameInfo.Level)},
		{"Health: ", fmt.Sprintf("%d/%d", v.playground.CharacterDTO.CurrentHealth, v.playground.CharacterDTO.MaxHealth)},
		{"Agile: ", strconv.Itoa(v.playground.CharacterDTO.Agile)},
		{"Strength: ", strconv.Itoa(v.playground.CharacterDTO.Strength)},
	}

	x := 4
	y := constants.MapHeight + OffsetCommonGameplay + 2

	for _, item := range items {
		view_utils.PrintText(v.screen, x, y, item.name, v.styles.Ui.Text)
		view_utils.PrintText(v.screen, x+len(item.name), y, item.value, v.styles.Ui.Value)
		x += len(item.name) + len(item.value) + 10
	}
}
