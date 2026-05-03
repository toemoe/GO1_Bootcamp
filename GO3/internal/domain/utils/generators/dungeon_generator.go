package generators

import (
	"container/list"
	"math/rand"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	"slices"

	"go.uber.org/zap"
)

type ConnectionNode struct {
	Nodes [4]*ConnectionNode
}

func (cn *ConnectionNode) SetConnection(node *ConnectionNode, dir value_objects.Directions) {
	cn.Nodes[dir] = node
	if node != nil {
		node.Nodes[value_objects.GetOppositeDirection(dir)] = cn
	}
}

func (cn *ConnectionNode) RemoveConnection(dir value_objects.Directions) {
	node := cn.Nodes[dir]
	if node != nil {
		cn.SetConnection(nil, dir)
		node.SetConnection(nil, value_objects.GetOppositeDirection(dir))
	}
}

type DungeonGenerator interface {
	GenerateDungeon() *entities.Dungeon
}

type dungeonGeneratorImpl struct {
	log utils.Logger
}

func NewDungeonGenerator(log utils.Logger) DungeonGenerator {
	return &dungeonGeneratorImpl{log: log}
}

func (dg *dungeonGeneratorImpl) GenerateDungeon() *entities.Dungeon {
	dg.log.Info("Generate dungeon")
	rooms := dg.generateRoomsWalls()
	connections := dg.createRoomsConnections()

	dg.createRoomsDoors(rooms, connections)
	corridors := dg.generateRoomsCorridors(rooms, connections)
	for dg.isCrossCorridors(corridors) {
		rooms = dg.generateRoomsWalls()
		connections = dg.createRoomsConnections()

		dg.createRoomsDoors(rooms, connections)
		corridors = dg.generateRoomsCorridors(rooms, connections)
	}
	portal := dg.generatePortal(rooms)

	return entities.NewDungeon(rooms, corridors, portal)
}

func (dg *dungeonGeneratorImpl) generateRoomsWalls() *[constants.RoomsPerSide][constants.RoomsPerSide]entities.Room {
	dg.log.Info("Generate rooms")
	rooms := [constants.RoomsPerSide][constants.RoomsPerSide]entities.Room{}
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			room := dg.generateRoom(j, i)
			rooms[i][j] = *room
		}
	}

	return &rooms
}

func (dg *dungeonGeneratorImpl) generateRoom(sectorOffsetX, sectorOffsetY int) *entities.Room {
	startCoordX := rand.Intn(constants.MaxRoomWidth - constants.MinRoomWidth)
	endCoordX := rand.Intn(constants.MaxRoomWidth)
	if (startCoordX + constants.MinRoomWidth) > endCoordX {
		endCoordX = startCoordX + constants.MinRoomWidth
	}

	startCoordY := rand.Intn(constants.MaxRoomHeight - constants.MinRoomHeight)
	endCoordY := rand.Intn(constants.MaxRoomHeight)
	if (startCoordY + constants.MinRoomHeight) > endCoordY {
		endCoordY = startCoordY + constants.MinRoomHeight
	}

	for range sectorOffsetX {
		startCoordX += (constants.MaxRoomWidth + constants.OffsetPerSector)
		endCoordX += (constants.MaxRoomWidth + constants.OffsetPerSector)
	}

	for range sectorOffsetY {
		startCoordY += (constants.MaxRoomHeight + constants.OffsetPerSector)
		endCoordY += (constants.MaxRoomHeight + constants.OffsetPerSector)
	}

	topLeft := *value_objects.NewPosition(startCoordX, endCoordY)
	botRight := *value_objects.NewPosition(endCoordX, startCoordY)

	res := entities.NewRoom(topLeft, botRight)
	return res
}

func (dg *dungeonGeneratorImpl) createRoomsConnections() *[constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode {
	nodes := [constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode{}
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if j != constants.RoomsPerSide-1 {
				nodes[i][j].SetConnection(&nodes[i][j+1], value_objects.Right)
			}
			if i != constants.RoomsPerSide-1 {
				nodes[i][j].SetConnection(&nodes[i+1][j], value_objects.Top)
			}
		}
	}

	countDelNodes := rand.Intn(constants.MaxDeletedRooms-constants.MinDeletedRooms) + constants.MinDeletedRooms
	for countDelNodes != 0 {
		if dg.deleteRandomNode(&nodes) {
			countDelNodes--
		}
	}
	return &nodes
}

func (dg *dungeonGeneratorImpl) createRoomsDoors(rooms *[constants.RoomsPerSide][constants.RoomsPerSide]entities.Room, nodes *[constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode) {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			dg.createDoors(&rooms[i][j], &nodes[i][j])
		}
	}
}

func (dg *dungeonGeneratorImpl) createDoors(room *entities.Room, node *ConnectionNode) {
	if node.Nodes[value_objects.Top] != nil {
		coordX := rand.Intn(room.BotRight.X-room.TopLeft.X-1) + room.TopLeft.X + 1
		coordY := room.TopLeft.Y + 1
		err := room.SetDoor(coordX, coordY)
		if err != nil {
			panic(err)
		}
	}
	if node.Nodes[value_objects.Bottom] != nil {
		coordX := rand.Intn(room.BotRight.X-room.TopLeft.X-1) + room.TopLeft.X + 1
		coordY := room.BotRight.Y - 1
		err := room.SetDoor(coordX, coordY)
		if err != nil {
			panic(err)
		}
	}

	if node.Nodes[value_objects.Right] != nil {
		coordX := room.BotRight.X + 1
		coordY := rand.Intn(room.TopLeft.Y-room.BotRight.Y-1) + room.BotRight.Y + 1
		err := room.SetDoor(coordX, coordY)
		if err != nil {
			panic(err)
		}
	}

	if node.Nodes[value_objects.Left] != nil {
		coordX := room.TopLeft.X - 1
		coordY := rand.Intn(room.TopLeft.Y-room.BotRight.Y-1) + room.BotRight.Y + 1
		err := room.SetDoor(coordX, coordY)
		if err != nil {
			panic(err)
		}
	}
}

func (dg *dungeonGeneratorImpl) generateRoomsCorridors(rooms *[constants.RoomsPerSide][constants.RoomsPerSide]entities.Room, nodes *[constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode) *[]entities.Corridor {
	corridors := make([]entities.Corridor, 0, constants.MaxCorridorCounts)
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if j != constants.RoomsPerSide-1 && nodes[i][j].Nodes[value_objects.Right] != nil {
				leftCoordinate := rooms[i][j].Doors[value_objects.Right]
				rightCoordinate := rooms[i][j+1].Doors[value_objects.Left]

				dg.log.Info("Create corridor with door coordinates",
					zap.Int("leftCoord X", leftCoordinate.X),
					zap.Int("leftCoord Y", leftCoordinate.Y),
					zap.Int("rightCoord X", rightCoordinate.X),
					zap.Int("rightCoord Y", rightCoordinate.Y))

				if leftCoordinate.Y == rightCoordinate.Y {
					corridor := entities.NewCorridor()
					corridor.Positions[0] = *value_objects.NewPosition(leftCoordinate.X, leftCoordinate.Y)
					corridor.Positions[3] = *value_objects.NewPosition(rightCoordinate.X, rightCoordinate.Y)
					corridors = append(corridors, *corridor)
				} else {
					var centerX int
					if rightCoordinate.X-leftCoordinate.X <= 2 {
						centerX = rightCoordinate.X - 1
					} else {
						centerX = rand.Intn(rightCoordinate.X-leftCoordinate.X-2) + leftCoordinate.X + 1
					}

					corridor := entities.NewCorridor()
					corridor.Positions[0] = *value_objects.NewPosition(leftCoordinate.X, leftCoordinate.Y)
					corridor.Positions[1] = *value_objects.NewPosition(centerX, leftCoordinate.Y)
					corridor.Positions[2] = *value_objects.NewPosition(centerX, rightCoordinate.Y)
					corridor.Positions[3] = *value_objects.NewPosition(rightCoordinate.X, rightCoordinate.Y)
					corridors = append(corridors, *corridor)
				}
			}
			if i != constants.RoomsPerSide-1 && nodes[i][j].Nodes[value_objects.Top] != nil {
				botCoordinate := rooms[i][j].Doors[value_objects.Top]
				topCoordinate := rooms[i+1][j].Doors[value_objects.Bottom]

				dg.log.Info("Create corridor with door coordinates",
					zap.Int("botCoord X", botCoordinate.X),
					zap.Int("botCoord Y", botCoordinate.Y),
					zap.Int("topCoord X", topCoordinate.X),
					zap.Int("topCoord Y", topCoordinate.Y))

				if botCoordinate.X == topCoordinate.X {

					corridor := entities.NewCorridor()
					corridor.Positions[0] = *value_objects.NewPosition(botCoordinate.X, botCoordinate.Y)
					corridor.Positions[3] = *value_objects.NewPosition(topCoordinate.X, topCoordinate.Y)
					corridors = append(corridors, *corridor)
				} else {
					var centerY int
					if topCoordinate.Y-botCoordinate.Y <= 2 {
						centerY = topCoordinate.Y - 1
					} else {
						centerY = rand.Intn(topCoordinate.Y-botCoordinate.Y-2) + botCoordinate.Y + 1
					}

					corridor := entities.NewCorridor()
					corridor.Positions[0] = *value_objects.NewPosition(botCoordinate.X, botCoordinate.Y)
					corridor.Positions[1] = *value_objects.NewPosition(botCoordinate.X, centerY)
					corridor.Positions[2] = *value_objects.NewPosition(topCoordinate.X, centerY)
					corridor.Positions[3] = *value_objects.NewPosition(topCoordinate.X, topCoordinate.Y)
					corridors = append(corridors, *corridor)
				}
			}
		}
	}
	return &corridors
}

func (dg *dungeonGeneratorImpl) deleteRandomNode(nodes *[constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode) bool {
	i := rand.Intn(constants.RoomsPerSide)
	j := rand.Intn(constants.RoomsPerSide)

	directions := make([]value_objects.Directions, 0, 4)
	for k := value_objects.Top; k <= value_objects.Left; k++ {
		if nodes[i][j].Nodes[k] != nil {
			directions = append(directions, k)
		}
	}

	if lenDir := len(directions); lenDir <= 1 {
		return false
	}

	index := rand.Intn(len(directions))
	deletedConnections := nodes[i][j].Nodes[directions[index]]
	nodes[i][j].RemoveConnection(directions[index])
	if dg.getUnconnectedNodes(nodes).Len() != 0 {
		nodes[i][j].SetConnection(deletedConnections, directions[index])
		return false
	}
	return true
}

func (dg *dungeonGeneratorImpl) getUnconnectedNodes(nodes *[constants.RoomsPerSide][constants.RoomsPerSide]ConnectionNode) *list.List {
	nodesMap := make(map[*ConnectionNode]bool)

	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			nodesMap[&nodes[i][j]] = false
		}
	}
	countChecked := 0
	bfs(&nodes[0][0], &nodesMap, &countChecked)

	res := list.New()

	if countChecked != (constants.RoomsPerSide * constants.RoomsPerSide) {
		for key := range nodesMap {
			if !nodesMap[key] {
				res.PushBack(key)
			}
		}
	}
	dg.log.Info("Count of connected rooms", zap.Int("connectedRooms", (constants.RoomsPerSide*constants.RoomsPerSide)-res.Len()))
	return res
}

func bfs(start *ConnectionNode, checkedNodes *map[*ConnectionNode]bool, countChecked *int) {
	(*checkedNodes)[start] = true
	*countChecked += 1
	if *countChecked == (constants.RoomsPerSide * constants.RoomsPerSide) {
		return
	}

	for i := value_objects.Top; i <= value_objects.Left; i++ {
		nextNode := start.Nodes[i]
		if nextNode != nil && !(*checkedNodes)[nextNode] {
			bfs(nextNode, checkedNodes, countChecked)
		}
	}
}

func (dg *dungeonGeneratorImpl) generatePortal(rooms *[constants.RoomsPerSide][constants.RoomsPerSide]entities.Room) *value_objects.Position {
	i := rand.Intn(constants.RoomsPerSide)
	j := rand.Intn(constants.RoomsPerSide)

	portalRoom := rooms[i][j]

	portalX := rand.Intn(portalRoom.BotRight.X-portalRoom.TopLeft.X) + portalRoom.TopLeft.X
	portalY := rand.Intn(portalRoom.TopLeft.Y-portalRoom.BotRight.Y) + portalRoom.BotRight.Y

	return &value_objects.Position{X: portalX, Y: portalY}
}

func (dg *dungeonGeneratorImpl) isCrossCorridors(corridors *[]entities.Corridor) bool {
	existedCoord := make(map[value_objects.Position]struct{})

	for _, c := range *corridors {
		coords := c.GetAllCoordinates()
		dilateCoords := getDilateCoords(coords)
		for _, p := range coords {
			_, isExist := existedCoord[*p]
			if isExist {
				return true
			}
			_, isExist = dilateCoords[*p]
			if isExist {
				return true
			}

			existedCoord[*p] = struct{}{}
		}
	}
	return false
}

/*
 ***
 * *
 * *
 * ***
 *   *
 *** *
   * *
   * *
   ***
*/
func getDilateCoords(coords []*value_objects.Position) map[value_objects.Position]struct{} {
	res := make(map[value_objects.Position]struct{})
	coordsNotPoints := make([]value_objects.Position, 0, len(coords))

	for _, c := range coords {
		coordsNotPoints = append(coordsNotPoints, *c)
	}

	for _, c := range coordsNotPoints {
		for x := -1; x < 2; x += 2 {
			for y := -1; y < 2; y += 2 {
				pos := *value_objects.NewPosition(c.X+x, c.Y+y)
				if !slices.Contains(coordsNotPoints, pos) {
					res[pos] = struct{}{}
				}
			}
		}
	}

	return res
}
