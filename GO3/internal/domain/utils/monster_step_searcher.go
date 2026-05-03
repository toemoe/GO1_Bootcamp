package utils

import (
	"container/list"
	"math"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"
	"sort"
)

type monsterDist struct {
	Monster entities.Monster
	Dist    float64
}

type monstersByDistance []monsterDist

func (mds monstersByDistance) Len() int           { return len(mds) }
func (mds monstersByDistance) Swap(i, j int)      { mds[i], mds[j] = mds[j], mds[i] }
func (mds monstersByDistance) Less(i, j int) bool { return mds[i].Dist < mds[j].Dist }

type roomDepth struct {
	i, j  int
	depth int
}

type MonsterPos struct {
	Monster entities.Monster
	Pos     *value_objects.Position
}

func GetSortedMonsterByStepPriority(monsters *list.List, character *entities.Character) []entities.Monster {
	var sortedList monstersByDistance
	for e := monsters.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		dist := math.Sqrt(math.Pow(float64(monster.GetPosition().X-character.Position.X), 2) + math.Pow(float64(monster.GetPosition().Y-character.Position.Y), 2))
		sortedList = append(sortedList, monsterDist{Monster: monster, Dist: dist})
	}
	sort.Sort(sortedList)

	result := make([]entities.Monster, len(sortedList))
	for i, md := range sortedList {
		result[i] = md.Monster
	}
	return result
}

func GetNextMonsterPosition(dungeon *entities.Dungeon, character *entities.Character, monster entities.Monster) *value_objects.Position {
	for _, c := range dungeon.Corridors {
		if c.InCorridor(monster.GetPosition().X, monster.GetPosition().Y) {
			return getNextMonsterPositionCorridor(dungeon, character, monster, &c)
		}
	}

	i, j, err := dungeon.SearchRoomIndexByPos(monster.GetPosition().X, monster.GetPosition().Y)
	if err != nil {
		panic("getNextMonsterPosition")
	}
	return getNextMonsterPositionRoom(monster, i, j, dungeon, character)
}

func getNextMonsterPositionCorridor(dungeon *entities.Dungeon, character *entities.Character, monster entities.Monster, c *entities.Corridor) *value_objects.Position {
	if c.InCorridor(character.Position.X, character.Position.Y) {
		return getNextMonsterPositionInSameCorridor(dungeon, character, monster, c)
	} else {
		return getNextMonsterPositionOuterCorridor(dungeon, character, monster, c)
	}
}

func getNextMonsterPositionRoom(monster entities.Monster, iRoom, jRoom int, dungeon *entities.Dungeon, character *entities.Character) *value_objects.Position {
	if dungeon.Rooms[iRoom][jRoom].InRoom(character.Position.X, character.Position.Y) {
		return getNextMonsterPositionInSameRoom(monster, character.Position)
	} else {
		return getNextMonsterPositionOuterRoom(monster, iRoom, jRoom, dungeon, character)
	}
}

func getNextMonsterPositionOuterRoom(monster entities.Monster, iRoom, jRoom int, dungeon *entities.Dungeon, character *entities.Character) *value_objects.Position {
	isFound, pos := getNextPositionSorroundingCorridor(monster, iRoom, jRoom, dungeon, character)
	if isFound {
		return pos
	}

	room := dungeon.Rooms[iRoom][jRoom]
	minDepth := 5
	var depth int
	var curPosition *value_objects.Position

	dir := []value_objects.Directions{value_objects.Top, value_objects.Right, value_objects.Bottom, value_objects.Left}
	appendX := []int{0, -1, 0, 1}
	appendY := []int{-1, 0, 1, 0}

	for i := range dir {
		corridor := dungeon.GetRoomCorridor(iRoom, jRoom, dir[i])
		if corridor != nil && corridor.InCorridor(character.Position.X, character.Position.Y) {
			preDoorPosition := value_objects.NewPosition(room.Doors[dir[i]].X+appendX[i], room.Doors[dir[i]].Y+appendY[i])
			isFound, depth = bfs(iRoom-appendY[i], jRoom-appendX[i], dungeon, character.Position)
			if isFound && depth < minDepth {
				minDepth = depth
				curPosition = preDoorPosition
			}
		}
	}

	if minDepth == 5 {
		panic("getNextMonsterPositionOuterRoom")
	}

	return getNextPosition(monster.GetPosition(), curPosition)
}

func getNextMonsterPositionInSameRoom(monster entities.Monster, characterPos *value_objects.Position) *value_objects.Position {
	return getNextPosition(monster.GetPosition(), characterPos)
}

func getNextMonsterPositionOuterCorridor(dungeon *entities.Dungeon, character *entities.Character, monster entities.Monster, c *entities.Corridor) *value_objects.Position {
	searchedDir := value_objects.Top
	iAppend := 1
	jAppend := 0
	if c.IsHorizontalCorridor() {
		searchedDir = value_objects.Right
		iAppend = 0
		jAppend = 1
	}

	firstRoomI, firstRoomJ := -1, -1
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if dungeon.Rooms[i][j].Doors[searchedDir].IsEqual(&c.Positions[0]) {
				firstRoomI, firstRoomJ = i, j
			}
		}
	}

	allCoordinates := c.GetAllCoordinates()

	if firstRoomI == -1 || firstRoomJ == -1 {
		panic("getNextMonsterPositionOuterCorridor")
	}

	_, firstDepth := bfs(firstRoomI, firstRoomJ, dungeon, character.Position)
	_, secondDepth := bfs(firstRoomI+iAppend, firstRoomJ+jAppend, dungeon, character.Position)

	indexMonster := c.GetIndexByPosition(monster.GetPosition())

	if indexMonster == -1 || (firstDepth == 5 && secondDepth == 5) {
		panic("getNextMonsterPositionOuterCorridor")
	}

	if secondDepth < firstDepth {
		return getNextCorridorPos(allCoordinates, indexMonster, true)
	}
	return getNextCorridorPos(allCoordinates, indexMonster, false)
}

func getNextMonsterPositionInSameCorridor(dungeon *entities.Dungeon, character *entities.Character, monster entities.Monster, c *entities.Corridor) *value_objects.Position {
	allCoordinates := c.GetAllCoordinates()
	indexCharacter := c.GetIndexByPosition(character.Position)
	indexMonster := c.GetIndexByPosition(monster.GetPosition())

	if indexCharacter == -1 || indexMonster == -1 {
		panic("this")
	}

	if indexCharacter > indexMonster {
		return getNextCorridorPos(allCoordinates, indexMonster, true)
	}
	return getNextCorridorPos(allCoordinates, indexMonster, false)
}

func getNextPositionSorroundingCorridor(monster entities.Monster, iRoom, jRoom int, dungeon *entities.Dungeon, character *entities.Character) (bool, *value_objects.Position) {
	room := dungeon.Rooms[iRoom][jRoom]

	dir := []value_objects.Directions{value_objects.Top, value_objects.Right, value_objects.Bottom, value_objects.Left}
	appendX := []int{0, -1, 0, 1}
	appendY := []int{-1, 0, 1, 0}

	for i := range dir {
		corridor := dungeon.GetRoomCorridor(iRoom, jRoom, dir[i])
		if corridor != nil && corridor.InCorridor(character.Position.X, character.Position.Y) {
			preDoorPosition := value_objects.NewPosition(room.Doors[dir[i]].X+appendX[i], room.Doors[dir[i]].Y+appendY[i])
			if monster.GetPosition().IsEqual(preDoorPosition) {
				return true, getNextPosition(monster.GetPosition(), &room.Doors[dir[i]])
			}
			return true, getNextPosition(monster.GetPosition(), preDoorPosition)
		}
	}

	return false, nil
}

func getNextCorridorPos(corridorPos []*value_objects.Position, currentIndex int, byDirection bool) *value_objects.Position {
	dir := 1
	if !byDirection {
		dir = -1
	}

	if byDirection && currentIndex != len(corridorPos)-1 {
		return corridorPos[currentIndex+1]
	}

	if !byDirection && currentIndex != 0 {
		return corridorPos[currentIndex-1]
	}

	// out of corridor
	curPos := corridorPos[currentIndex]
	prevPos := corridorPos[currentIndex-dir]
	if curPos.X == prevPos.X {
		return value_objects.NewPosition(curPos.X, curPos.Y+dir)
	}
	return value_objects.NewPosition(curPos.X+dir, curPos.Y)
}

func bfs(iRoom, jRoom int, dungeon *entities.Dungeon, foundPosition *value_objects.Position) (bool, int) {
	var visited [constants.RoomsPerSide][constants.RoomsPerSide]bool
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			visited[i][j] = false
		}
	}

	queue := list.New()
	queue.PushBack(roomDepth{iRoom, jRoom, 0})

	for queue.Len() != 0 {
		node := queue.Remove(queue.Front()).(roomDepth)
		visited[node.i][node.j] = true
		if node.depth > 2 {
			continue
		}

		if dungeon.Rooms[node.i][node.j].InRoom(foundPosition.X, foundPosition.Y) {
			return true, node.depth
		}

		dirs := []value_objects.Directions{value_objects.Top, value_objects.Right, value_objects.Bottom, value_objects.Left}
		iNodeAppend := []int{1, 0, -1, 0}
		jNodeAppend := []int{0, 1, 0, -1}

		for i := range dirs {
			if !dungeon.Rooms[node.i][node.j].Doors[dirs[i]].IsDefault() {
				corridor := dungeon.GetRoomCorridor(node.i, node.j, dirs[i])
				if corridor.InCorridor(foundPosition.X, foundPosition.Y) {
					return true, node.depth + 1
				}
				if !visited[node.i+iNodeAppend[i]][node.j+jNodeAppend[i]] {
					queue.PushBack(roomDepth{node.i + iNodeAppend[i], node.j + jNodeAppend[i], node.depth + 1})
				}
			}
		}
	}

	return false, 5
}

func getNextPosition(startPos, endPos *value_objects.Position) *value_objects.Position {
	deltaX := startPos.X - endPos.X
	if deltaX < 0 {
		deltaX = -deltaX
	}

	deltaY := startPos.Y - endPos.Y
	if deltaY < 0 {
		deltaY = -deltaY
	}

	if deltaX > deltaY {
		if endPos.X > startPos.X {
			return value_objects.NewPosition(startPos.X+1, startPos.Y)
		}
		return value_objects.NewPosition(startPos.X-1, startPos.Y)
	}

	if endPos.Y > startPos.Y {
		return value_objects.NewPosition(startPos.X, startPos.Y+1)
	}
	return value_objects.NewPosition(startPos.X, startPos.Y-1)
}
