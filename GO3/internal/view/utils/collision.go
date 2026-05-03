package utils

import (
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/value_objects"
	"slices"
)

func IsInRoom(position value_objects.Position, room *dto.RoomDTO) bool {
	if room == nil {
		return false
	}

	return position.X >= room.TopLeft.X && position.X <= room.BotRight.X &&
		position.Y >= room.BotRight.Y && position.Y <= room.TopLeft.Y
}

func IsInCorridor(pos value_objects.Position, c *dto.CorridorDTO) bool {
	isDirect := IsDirectCorridor(c)
	isHorizontal := IsHorizontalCorridor(c, isDirect)

	if isDirect {
		if isHorizontal {
			return pos.Y == c.Positions[0].Y &&
				c.Positions[0].X <= pos.X && pos.X <= c.Positions[3].X
		}
		return pos.X == c.Positions[0].X &&
			c.Positions[0].Y <= pos.Y && pos.Y <= c.Positions[3].Y
	}

	if isHorizontal {
		if (pos.Y == c.Positions[0].Y &&
			c.Positions[0].X <= pos.X && pos.X <= c.Positions[1].X) ||
			(pos.Y == c.Positions[2].Y && c.Positions[2].X <= pos.X && pos.X <= c.Positions[3].X) {
			return true
		}

		if c.Positions[1].Y < c.Positions[2].Y {
			return pos.X == c.Positions[1].X && c.Positions[1].Y <= pos.Y && pos.Y <= c.Positions[2].Y
		}

		return pos.X == c.Positions[1].X && c.Positions[2].Y <= pos.Y && pos.Y <= c.Positions[1].Y
	}

	if (pos.X == c.Positions[0].X && c.Positions[0].Y <= pos.Y && pos.Y <= c.Positions[1].Y) ||
		(pos.X == c.Positions[2].X && c.Positions[2].Y <= pos.Y && pos.Y <= c.Positions[3].Y) {
		return true
	}

	if c.Positions[1].X < c.Positions[2].X {
		return pos.Y == c.Positions[1].Y && c.Positions[1].X <= pos.X && pos.X <= c.Positions[2].X
	}

	return pos.Y == c.Positions[1].Y && c.Positions[2].X <= pos.X && pos.X <= c.Positions[1].X
}

func IsDirectCorridor(corridor *dto.CorridorDTO) bool {
	return corridor.Positions[1].X == value_objects.UninitializedCoord
}

func IsHorizontalCorridor(corridor *dto.CorridorDTO, isDirect bool) bool {
	compareIndex := 1
	if isDirect {
		compareIndex = 3

	}
	return corridor.Positions[0].Y == corridor.Positions[compareIndex].Y
}

func GetCurrentRoom(playground *dto.PlaygroundDTO) *dto.RoomDTO {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if playground.DungeonDTO.Rooms[i][j].VisitStateDTO == "current_space" {
				return &playground.DungeonDTO.Rooms[i][j]
			}
		}
	}
	return nil
}

func getAllCoordinates(corridor *dto.CorridorDTO) []*value_objects.Position {
	res := make([]*value_objects.Position, 0)
	isDirectCorridor := IsDirectCorridor(corridor)
	isHorizontalCorridor := IsHorizontalCorridor(corridor, isDirectCorridor)
	if isDirectCorridor && isHorizontalCorridor {
		for x := corridor.Positions[0].X; x <= corridor.Positions[3].X; x++ {
			res = append(res, value_objects.NewPosition(x, corridor.Positions[0].Y))
		}
		return res
	} else if isDirectCorridor && !isHorizontalCorridor {
		for y := corridor.Positions[0].Y; y <= corridor.Positions[3].Y; y++ {
			res = append(res, value_objects.NewPosition(corridor.Positions[0].X, y))
		}
		return res
	} else if !isDirectCorridor && isHorizontalCorridor {
		for x := corridor.Positions[0].X; x <= corridor.Positions[1].X; x++ {
			res = append(res, value_objects.NewPosition(x, corridor.Positions[0].Y))
		}
		if corridor.Positions[1].Y < corridor.Positions[2].Y {
			for y := corridor.Positions[1].Y + 1; y < corridor.Positions[2].Y; y++ {
				res = append(res, value_objects.NewPosition(corridor.Positions[1].X, y))
			}
		} else {
			for y := corridor.Positions[1].Y - 1; y > corridor.Positions[2].Y; y-- {
				res = append(res, value_objects.NewPosition(corridor.Positions[1].X, y))
			}
		}
		for x := corridor.Positions[2].X; x <= corridor.Positions[3].X; x++ {
			res = append(res, value_objects.NewPosition(x, corridor.Positions[2].Y))
		}
		return res
	} else {
		for y := corridor.Positions[0].Y; y <= corridor.Positions[1].Y; y++ {
			res = append(res, value_objects.NewPosition(corridor.Positions[0].X, y))
		}
		if corridor.Positions[1].X < corridor.Positions[2].X {
			for x := corridor.Positions[1].X + 1; x < corridor.Positions[2].X; x++ {
				res = append(res, value_objects.NewPosition(x, corridor.Positions[1].Y))
			}
		} else {
			for x := corridor.Positions[1].X - 1; x > corridor.Positions[2].X; x-- {
				res = append(res, value_objects.NewPosition(x, corridor.Positions[1].Y))
			}
		}
		for y := corridor.Positions[2].Y; y <= corridor.Positions[3].Y; y++ {
			res = append(res, value_objects.NewPosition(corridor.Positions[2].X, y))
		}
		return res
	}
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
func GetDilateCorridorCoords(corridor *dto.CorridorDTO) []*value_objects.Position {
	coords := getAllCoordinates(corridor)
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

	resSlice := make([]*value_objects.Position, 0, len(res))
	for k := range res {
		resSlice = append(resSlice, &k)
	}
	return resSlice
}
