package entities

import (
	"fmt"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/value_objects"
)

type Dungeon struct {
	Rooms     [constants.RoomsPerSide][constants.RoomsPerSide]Room `json:"rooms" validate:"dive,dive"`
	Corridors []Corridor                                           `json:"corridors"`
	Portal    *value_objects.Position                              `json:"portal"`
}

func NewDungeon(rooms *[constants.RoomsPerSide][constants.RoomsPerSide]Room, corridors *[]Corridor, portal *value_objects.Position) *Dungeon {
	return &Dungeon{Rooms: *rooms, Corridors: *corridors, Portal: portal}
}

func (d *Dungeon) InDungeon(pos *value_objects.Position) bool {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if d.Rooms[i][j].InRoom(pos.X, pos.Y) {
				return true
			}
		}
	}

	for _, c := range d.Corridors {
		if c.InCorridor(pos.X, pos.Y) {
			return true
		}
	}
	return false
}

func (d *Dungeon) InCorridors(pos *value_objects.Position) bool {
	for _, c := range d.Corridors {
		if c.InCorridor(pos.X, pos.Y) {
			return true
		}
	}
	return false
}

func (d *Dungeon) InRooms(pos *value_objects.Position) bool {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if d.Rooms[i][j].InRoom(pos.X, pos.Y) {
				return true
			}
		}
	}
	return false
}

func (d *Dungeon) GetRoomCorridor(i, j int, dir value_objects.Directions) *Corridor {
	roomDoor := d.Rooms[i][j].Doors[dir]
	if roomDoor.IsDefault() {
		return nil
	}

	checkedIndex := 3
	if dir == value_objects.Top || dir == value_objects.Right {
		checkedIndex = 0
	}

	for i := range d.Corridors {
		if d.Corridors[i].Positions[checkedIndex].IsEqual(&roomDoor) {
			return &d.Corridors[i]
		}
	}
	panic("dungeon Corridor")
}

func (d *Dungeon) SearchRoomIndexByPos(x, y int) (int, int, error) {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if d.Rooms[i][j].InRoom(x, y) {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("Index Not Found")
}

func (d *Dungeon) UpdateVisitedState(curPosition *value_objects.Position) {
	i, j, err := d.SearchRoomIndexByPos(curPosition.X, curPosition.Y)
	if err == nil {
		d.clearCurrentSpaceVisitState()
		d.Rooms[i][j].VisitState = value_objects.CurrentSpace
		return
	}

	for i := range d.Corridors {
		if d.Corridors[i].InCorridor(curPosition.X, curPosition.Y) {
			d.clearCurrentSpaceVisitState()
			d.Corridors[i].VisitState = value_objects.CurrentSpace
		}
	}
}

func (d *Dungeon) clearCurrentSpaceVisitState() {
	for i := range constants.RoomsPerSide {
		for j := range constants.RoomsPerSide {
			if d.Rooms[i][j].VisitState == value_objects.CurrentSpace {
				d.Rooms[i][j].VisitState = value_objects.VisitedState
			}
		}
	}
	for i := range d.Corridors {
		if d.Corridors[i].VisitState == value_objects.CurrentSpace {
			d.Corridors[i].VisitState = value_objects.VisitedState
		}
	}
}
