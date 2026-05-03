package services

import (
	"container/list"
	"math/rand"
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/domain/aggregates"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/utils/generators"
	"s21_rogue/internal/domain/value_objects"

	"github.com/google/uuid"
)

type PlaygroundGeneratorService interface {
	GenerateNewLevel(*aggregates.Playground, int) *aggregates.Playground
}

func NewPlaygroundGeneratorService(generator generators.DungeonGenerator, log utils.Logger) PlaygroundGeneratorService {
	return &playgroundGeneratorServiceImpl{generator: generator, log: log}

}

type playgroundGeneratorServiceImpl struct {
	generator generators.DungeonGenerator
	log       utils.Logger
}

func (pg *playgroundGeneratorServiceImpl) GenerateNewLevel(p *aggregates.Playground, level int) *aggregates.Playground {
	excludePos := make(map[value_objects.Position]struct{})

	dungeon := pg.generator.GenerateDungeon()

	portalRoomI, portalRoomJ, _ := dungeon.SearchRoomIndexByPos(dungeon.Portal.X, dungeon.Portal.Y)
	excludePos[*dungeon.Portal] = struct{}{}

	characterPos := generateCoordinateInOtherRoom(dungeon, portalRoomI, portalRoomJ, &excludePos)
	excludePos[*characterPos] = struct{}{}

	mainRoomI, mainRoomJ, _ := dungeon.SearchRoomIndexByPos(characterPos.X, characterPos.Y)
	dungeon.Rooms[mainRoomI][mainRoomJ].VisitState = value_objects.CurrentSpace

	monstersList := pg.generateMonsters(dungeon, mainRoomI, mainRoomJ, level, &excludePos)

	countItems := rand.Intn(4) + max(0, 10-level)
	var countItemList [4]int
	for range countItems {
		index := rand.Intn(4)
		countItemList[index]++
	}
	positions := make([]*value_objects.Position, 0, countItems)
	for range countItems {
		newPos := generateCoordinateInOtherRoom(dungeon, mainRoomI, mainRoomJ, &excludePos)
		positions = append(positions, newPos)
		excludePos[*newPos] = struct{}{}
	}
	indexItem := 0
	indexPosition := 0

	foods := make(map[value_objects.Position]entities.Food)
	for range countItemList[indexItem] {
		foods[*positions[indexPosition]] = entities.GenerateFood()
		indexPosition++
	}
	indexItem++

	scrolls := make(map[value_objects.Position]entities.Scroll)
	for range countItemList[indexItem] {
		scrolls[*positions[indexPosition]] = entities.GenerateScroll()
		indexPosition++
	}
	indexItem++

	elixirs := make(map[value_objects.Position]entities.Elixir)
	for range countItemList[indexItem] {
		elixirs[*positions[indexPosition]] = entities.GenerateElixir()
		indexPosition++
	}
	indexItem++

	weapons := make(map[value_objects.Position]*entities.Weapon)
	for range countItemList[indexItem] {
		weapons[*positions[indexPosition]] = entities.GenerateWeapon()
		indexPosition++
	}

	id := uuid.New()
	character := entities.NewCharacter(characterPos, 100, 100, 25, 10)
	backpack := entities.NewBackpack()

	if p != nil {
		id = p.Id
		character, _ = mappers.MapFromCharacterDTO(mappers.MapToCharacterDTO(p.Character))
		character.SetPosition(characterPos)
		backpack = p.Backpack
	}

	return aggregates.NewPlayground(id, dungeon, character, backpack, monstersList, foods, scrolls, elixirs, weapons)
}

func (pg *playgroundGeneratorServiceImpl) generateMonsters(
	dungeon *entities.Dungeon,
	excludeRoomI,
	excludeRoomJ int,
	level int,
	excludePos *map[value_objects.Position]struct{}) *list.List { // Monster

	monstersList := list.New()
	countMonsters := (3 + level) / 2

	maxMonsterType := entities.VampireType
	if level > 2 {
		maxMonsterType++
	}
	if level > 5 {
		maxMonsterType++
	}
	if level > 9 {
		maxMonsterType++
	}

	for range countMonsters {
		monsterPos := generateCoordinateInOtherRoom(dungeon, excludeRoomI, excludeRoomJ, excludePos)
		monsterType := rand.Intn(int(maxMonsterType) + 1)
		monstersList.PushFront(entities.NewMonster(monsterPos, entities.MonsterType(monsterType)))
		(*excludePos)[*monsterPos] = struct{}{}
	}

	return monstersList
}

func generateCoordinateInOtherRoom(dungeon *entities.Dungeon, excludeRoomI, excludeRoomJ int, excludePos *map[value_objects.Position]struct{}) *value_objects.Position {
	i, j := rand.Intn(constants.RoomsPerSide), rand.Intn(constants.RoomsPerSide)
	for i == excludeRoomI && j == excludeRoomJ {
		i, j = rand.Intn(constants.RoomsPerSide), rand.Intn(constants.RoomsPerSide)
	}

	targetRoom := dungeon.Rooms[i][j]
	for {
		targetX := rand.Intn(targetRoom.BotRight.X-targetRoom.TopLeft.X) + targetRoom.TopLeft.X
		targetY := rand.Intn(targetRoom.TopLeft.Y-targetRoom.BotRight.Y) + targetRoom.BotRight.Y
		targetPos := value_objects.NewPosition(targetX, targetY)
		_, isBusy := (*excludePos)[*targetPos]
		if !isBusy {
			return targetPos
		}
	}

}
