package validators

import (
	"container/list"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"

	"github.com/go-playground/validator/v10"
)

func DungeonStructLevelValidation(sl validator.StructLevel) {
	dungeon := sl.Current().Interface().(entities.Dungeon)
	if !validateRoomSize(&dungeon) {
		sl.ReportError(dungeon.Rooms, "Room", "RoomSize", "roomsizeleft", "")
		return
	}

	if !validateRoomDoors(&dungeon) {
		sl.ReportError(dungeon, "Door", "Doors", "doors", "")
		return
	}

	if !validateCorridors(&dungeon) {
		sl.ReportError(dungeon.Corridors, "Corridors", "Corridors", "corridor", "")
		return
	}

}

func validateRoomSize(dungeon *entities.Dungeon) bool {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if dungeon.Rooms[i][j].TopLeft.X < (j * (constants.OffsetPerSector + constants.MaxRoomWidth)) {
				return false
			}

			if dungeon.Rooms[i][j].BotRight.X > (j*(constants.OffsetPerSector+constants.MaxRoomWidth) + constants.MaxRoomWidth) {
				return false
			}

			if dungeon.Rooms[i][j].BotRight.Y < (i * (constants.OffsetPerSector + constants.MaxRoomHeight)) {
				return false
			}

			if dungeon.Rooms[i][j].TopLeft.Y > (i*(constants.OffsetPerSector+constants.MaxRoomHeight) + constants.MaxRoomHeight) {
				return false
			}

		}
	}
	return true
}

func validateRoomDoors(dungeon *entities.Dungeon) bool {
	// outer Doors
	for i := range constants.RoomsPerSide {
		if !dungeon.Rooms[i][0].Doors[value_objects.Left].IsDefault() {
			return false
		}

		if !dungeon.Rooms[0][i].Doors[value_objects.Bottom].IsDefault() {
			return false
		}

		if !dungeon.Rooms[i][constants.RoomsPerSide-1].Doors[value_objects.Right].IsDefault() {
			return false
		}

		if !dungeon.Rooms[constants.RoomsPerSide-1][i].Doors[value_objects.Top].IsDefault() {
			return false
		}
	}

	// inner Doors correct connections

	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if !dungeon.Rooms[i][j].Doors[value_objects.Right].IsDefault() &&
				dungeon.Rooms[i][j+1].Doors[value_objects.Left].IsDefault() {
				return false
			}

			if !dungeon.Rooms[i][j].Doors[value_objects.Top].IsDefault() &&
				dungeon.Rooms[i+1][j].Doors[value_objects.Bottom].IsDefault() {
				return false
			}

			if j == 1 && !dungeon.Rooms[i][j].Doors[value_objects.Left].IsDefault() &&
				dungeon.Rooms[i][j-1].Doors[value_objects.Right].IsDefault() {
				return false
			}

			if i == 1 && !dungeon.Rooms[i][j].Doors[value_objects.Bottom].IsDefault() &&
				dungeon.Rooms[i-1][j].Doors[value_objects.Top].IsDefault() {
				return false
			}
		}
	}
	return true
}

func validateCorridors(dungeon *entities.Dungeon) bool {
	valid, estHorCountCorridors := validateHorizontalCorridors(dungeon)
	if !valid {
		return false
	}

	valid, estVertCountCorridors := validateVerticalCorridors(dungeon)
	if !valid {
		return false
	}

	if (estHorCountCorridors + estVertCountCorridors) != len(dungeon.Corridors) {
		return false
	}
	return true
}

func validateVerticalCorridors(dungeon *entities.Dungeon) (bool, int) {
	verticalCorridors := list.New()

	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if i != constants.RoomsPerSide-1 && !dungeon.Rooms[i][j].Doors[value_objects.Top].IsDefault() {
				start := dungeon.Rooms[i][j].Doors[value_objects.Top]
				end := dungeon.Rooms[i+1][j].Doors[value_objects.Bottom]

				verticalCorridors.PushBack([2]value_objects.Position{{X: start.X, Y: start.Y}, {X: end.X, Y: end.Y}})
			}
		}
	}

	for c := verticalCorridors.Front(); c != nil; c = c.Next() {
		isFound := false
		for i := 0; i < len(dungeon.Corridors) && !isFound; i++ {
			estimatedCorridor := c.Value.([2]value_objects.Position)
			if dungeon.Corridors[i].Positions[0].X == estimatedCorridor[0].X &&
				dungeon.Corridors[i].Positions[0].Y == estimatedCorridor[0].Y &&
				dungeon.Corridors[i].Positions[3].X == estimatedCorridor[1].X &&
				dungeon.Corridors[i].Positions[3].Y == estimatedCorridor[1].Y {
				if dungeon.Corridors[i].Positions[1].IsDefault() && dungeon.Corridors[i].Positions[2].IsDefault() {
					isFound = true
					break
				} else if dungeon.Corridors[i].Positions[1].Y == dungeon.Corridors[i].Positions[2].Y &&
					dungeon.Corridors[i].Positions[1].X == estimatedCorridor[0].X &&
					dungeon.Corridors[i].Positions[2].X == estimatedCorridor[1].X {
					isFound = true
				}
			}
		}
		if !isFound {
			return false, verticalCorridors.Len()
		}

	}

	return true, verticalCorridors.Len()
}

func validateHorizontalCorridors(dungeon *entities.Dungeon) (bool, int) {
	horizontalCorridors := list.New()

	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if j != constants.RoomsPerSide-1 && !dungeon.Rooms[i][j].Doors[value_objects.Right].IsDefault() {
				start := dungeon.Rooms[i][j].Doors[value_objects.Right]
				end := dungeon.Rooms[i][j+1].Doors[value_objects.Left]

				horizontalCorridors.PushBack([2]value_objects.Position{{X: start.X, Y: start.Y}, {X: end.X, Y: end.Y}})
			}
		}
	}

	for c := horizontalCorridors.Front(); c != nil; c = c.Next() {
		isFound := false
		for i := 0; i < len(dungeon.Corridors) && !isFound; i++ {
			estimatedCorridor := c.Value.([2]value_objects.Position)
			if dungeon.Corridors[i].Positions[0].X == estimatedCorridor[0].X &&
				dungeon.Corridors[i].Positions[0].Y == estimatedCorridor[0].Y &&
				dungeon.Corridors[i].Positions[3].X == estimatedCorridor[1].X &&
				dungeon.Corridors[i].Positions[3].Y == estimatedCorridor[1].Y {
				if dungeon.Corridors[i].Positions[1].IsDefault() && dungeon.Corridors[i].Positions[2].IsDefault() {
					isFound = true
					break
				} else if dungeon.Corridors[i].Positions[1].X == dungeon.Corridors[i].Positions[2].X &&
					dungeon.Corridors[i].Positions[1].Y == estimatedCorridor[0].Y &&
					dungeon.Corridors[i].Positions[2].Y == estimatedCorridor[1].Y {
					isFound = true
				}
			}
		}
		if !isFound {
			return false, horizontalCorridors.Len()
		}

	}
	return true, horizontalCorridors.Len()
}

func getDirectionCorridorsList(dungeon *entities.Dungeon, dir value_objects.Directions) *list.List {
	corridorList := list.New()

	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if j != constants.RoomsPerSide-1 && !dungeon.Rooms[i][j].Doors[dir].IsDefault() {
				start := dungeon.Rooms[i][j].Doors[dir]
				var end value_objects.Position
				if dir == value_objects.Right {
					end = dungeon.Rooms[i][j+1].Doors[value_objects.GetOppositeDirection(dir)]
				} else {
					end = dungeon.Rooms[i+1][j].Doors[value_objects.GetOppositeDirection(dir)]
				}

				corridorList.PushBack([2]value_objects.Position{{X: start.X, Y: start.Y}, {X: end.X, Y: end.Y}})
			}
		}
	}
	return corridorList

}
