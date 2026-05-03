package dto

import (
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/value_objects"
)

type DungeonDTO struct {
	Rooms        [constants.RoomsPerSide][constants.RoomsPerSide]RoomDTO `json:"rooms"`
	CorridorsDTO []CorridorDTO                                           `json:"corridors"`
	Portal       value_objects.Position                                  `json:"portal"`
}

type RoomDTO struct {
	TopLeft       value_objects.Position    `json:"top_left"`
	BotRight      value_objects.Position    `json:"bot_right"`
	Doors         [4]value_objects.Position `json:"doors"`
	VisitStateDTO string                    `json:"visited_state"`
}

type CorridorDTO struct {
	Positions     [4]value_objects.Position `json:"position"`
	VisitStateDTO string                    `json:"visited_state"`
}
