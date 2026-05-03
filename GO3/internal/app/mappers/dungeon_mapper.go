package mappers

import (
	"fmt"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"
)

type VisitStateDTO string

const (
	VisitedStateDTO    string = "visited"
	NotVisitedStateDTO string = "not_visited"
	CurrentSpaceDTO    string = "current_space"
)

func MapToDungeonDTO(dungeon *entities.Dungeon) *dto.DungeonDTO {
	dungeonDTO := dto.DungeonDTO{}
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			dungeonDTO.Rooms[i][j] = *MapToRoomDTO(&dungeon.Rooms[i][j])
		}
	}
	dungeonDTO.Portal = *dungeon.Portal

	dungeonDTO.CorridorsDTO = make([]dto.CorridorDTO, len(dungeon.Corridors))
	for i := range dungeon.Corridors {
		dungeonDTO.CorridorsDTO[i] = *MapToCorridorDTO(&dungeon.Corridors[i])
	}
	return &dungeonDTO
}

func MapFromDungeonDTO(dungeonDTO *dto.DungeonDTO) (*entities.Dungeon, error) {
	corridors := make([]entities.Corridor, len(dungeonDTO.CorridorsDTO))
	for i := range dungeonDTO.CorridorsDTO {
		curCorridor, err := MapFromCorridorDTO(&dungeonDTO.CorridorsDTO[i])
		if err != nil {
			return nil, err
		}
		corridors[i] = *curCorridor
	}
	var rooms [constants.RoomsPerSide][constants.RoomsPerSide]entities.Room
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			curRoom, err := MapFromRoomDTO(&dungeonDTO.Rooms[i][j])
			if err != nil {
				return nil, err
			}
			rooms[i][j] = *curRoom
		}
	}

	return entities.NewDungeon(&rooms, &corridors, &dungeonDTO.Portal), nil
}

func MapToRoomDTO(room *entities.Room) *dto.RoomDTO {
	roomDTO := dto.RoomDTO{}
	roomDTO.TopLeft = room.TopLeft
	roomDTO.BotRight = room.BotRight
	for i := range 4 {
		roomDTO.Doors[i] = room.Doors[i]
	}

	switch room.VisitState {
	case value_objects.VisitedState:
		roomDTO.VisitStateDTO = VisitedStateDTO
	case value_objects.NotVisitedState:
		roomDTO.VisitStateDTO = NotVisitedStateDTO
	case value_objects.CurrentSpace:
		roomDTO.VisitStateDTO = CurrentSpaceDTO
	}

	return &roomDTO
}

func MapFromRoomDTO(roomDTO *dto.RoomDTO) (*entities.Room, error) {
	room := entities.NewRoom(roomDTO.TopLeft, roomDTO.BotRight)
	for i := range roomDTO.Doors {
		room.Doors[i] = roomDTO.Doors[i]
	}
	switch roomDTO.VisitStateDTO {
	case VisitedStateDTO:
		room.VisitState = value_objects.VisitedState
	case NotVisitedStateDTO:
		room.VisitState = value_objects.NotVisitedState
	case CurrentSpaceDTO:
		room.VisitState = value_objects.CurrentSpace
	default:
		return nil, fmt.Errorf("Incorrect parse visit state for roomDTO")
	}

	return room, nil
}

func MapToCorridorDTO(corridor *entities.Corridor) *dto.CorridorDTO {
	corridorDTO := dto.CorridorDTO{}
	for i := range 4 {
		corridorDTO.Positions[i] = corridor.Positions[i]
	}

	switch corridor.VisitState {
	case value_objects.VisitedState:
		corridorDTO.VisitStateDTO = VisitedStateDTO
	case value_objects.NotVisitedState:
		corridorDTO.VisitStateDTO = NotVisitedStateDTO
	case value_objects.CurrentSpace:
		corridorDTO.VisitStateDTO = CurrentSpaceDTO
	}

	return &corridorDTO
}

func MapFromCorridorDTO(corridorDTO *dto.CorridorDTO) (*entities.Corridor, error) {
	corridor := entities.NewCorridor()
	for i := range corridorDTO.Positions {
		corridor.Positions[i] = corridorDTO.Positions[i]
	}
	switch corridorDTO.VisitStateDTO {
	case VisitedStateDTO:
		corridor.VisitState = value_objects.VisitedState
	case NotVisitedStateDTO:
		corridor.VisitState = value_objects.NotVisitedState
	case CurrentSpaceDTO:
		corridor.VisitState = value_objects.CurrentSpace
	default:
		return nil, fmt.Errorf("Incorrect parse visit state for corridorDTO")
	}

	return corridor, nil
}
