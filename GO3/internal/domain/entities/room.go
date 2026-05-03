package entities

import (
	"fmt"
	"s21_rogue/internal/domain/value_objects"
)

type Room struct {
	TopLeft    value_objects.Position    `json:"top_left"`
	BotRight   value_objects.Position    `json:"bot_right"`
	Doors      [4]value_objects.Position `json:"doors"`
	VisitState value_objects.VisitState  `json:"visited_state"`
}

func NewRoom(topLeft, botRight value_objects.Position) *Room {
	doors := [4]value_objects.Position{}
	for i := range 4 {
		doors[i] = *value_objects.NewDefaultPosition()
	}

	return &Room{
		TopLeft:    topLeft,
		BotRight:   botRight,
		Doors:      doors,
		VisitState: value_objects.NotVisitedState,
	}
}

func (r *Room) SetDoor(x, y int) error {
	if x == r.TopLeft.X-1 && (r.BotRight.Y < y && y < r.TopLeft.Y) {
		r.Doors[value_objects.Left] = *value_objects.NewPosition(x, y)
	} else if x == r.BotRight.X+1 && (r.BotRight.Y < y && y < r.TopLeft.Y) {
		r.Doors[value_objects.Right] = *value_objects.NewPosition(x, y)
	} else if y == r.TopLeft.Y+1 && (r.TopLeft.X < x && x < r.BotRight.X) {
		r.Doors[value_objects.Top] = *value_objects.NewPosition(x, y)
	} else if y == r.BotRight.Y-1 && (r.TopLeft.X < x && x < r.BotRight.X) {
		r.Doors[value_objects.Bottom] = *value_objects.NewPosition(x, y)
	} else {
		return fmt.Errorf("coordinate door out of perimeter {%v, %v}", x, y)
	}
	return nil
}

func (r *Room) InRoom(x, y int) bool {
	return r.TopLeft.X <= x &&
		x <= r.BotRight.X &&
		r.BotRight.Y <= y &&
		y <= r.TopLeft.Y
}
