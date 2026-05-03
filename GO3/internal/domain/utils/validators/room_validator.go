package validators

import (
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"

	"github.com/go-playground/validator/v10"
)

func RoomStructLevelValidation(sl validator.StructLevel) {
	room := sl.Current().Interface().(entities.Room)

	if room.TopLeft.X < 0 || room.BotRight.Y < 0 {
		sl.ReportError(room.TopLeft.X, "top_left", "TopLeft.X", "tleftorbright", "")
		sl.ReportError(room.BotRight.Y, "bot_right", "BotRight.Y", "tleftorbright", "")
		return
	}

	if room.TopLeft.Y >= constants.MapHeight || room.BotRight.X >= constants.MapWidth {
		sl.ReportError(room.TopLeft.Y, "top_left", "TopLeft.Y", "tleftorbright", "")
		sl.ReportError(room.BotRight.X, "bot_right", "BotRight.X", "tleftorbright", "")
		return
	}

	if room.TopLeft.X+constants.MinRoomWidth > room.BotRight.X {
		sl.ReportError(room.TopLeft.X, "top_left", "TopLeft.X", "tleftorbright", "")
		sl.ReportError(room.BotRight.X, "bot_right", "BotRight.X", "tleftorbright", "")
		return
	}

	if (room.BotRight.X - room.TopLeft.X) >= constants.MaxRoomWidth {
		sl.ReportError(room.TopLeft.X, "top_left", "TopLeft.X", "tleftorbright", "")
		sl.ReportError(room.BotRight.X, "bot_right", "BotRight.X", "tleftorbright", "")
		return
	}

	if room.BotRight.Y+constants.MinRoomHeight > room.TopLeft.Y {
		sl.ReportError(room.TopLeft.Y, "top_left", "TopLeft.X", "tleftorbright", "")
		sl.ReportError(room.BotRight.Y, "bot_right", "BotRight.Y", "tleftorbright", "")
		return
	}

	if (room.TopLeft.Y - room.BotRight.Y) >= constants.MaxRoomHeight {
		sl.ReportError(room.TopLeft.Y, "top_left", "TopLeft.X", "tleftorbright", "")
		sl.ReportError(room.BotRight.Y, "bot_right", "BotRight.Y", "tleftorbright", "")
		return
	}

	x, y := room.Doors[value_objects.Left].X, room.Doors[value_objects.Left].Y
	if !(x == value_objects.UninitializedCoord && y == value_objects.UninitializedCoord) {
		if !(x == room.TopLeft.X-1 && (room.BotRight.Y < y && y < room.TopLeft.Y)) {
			sl.ReportError(room.Doors, "door", "room.Doors[Left]", "door", "")
			return
		}
	}

	x, y = room.Doors[value_objects.Right].X, room.Doors[value_objects.Right].Y
	if !(x == value_objects.UninitializedCoord && y == value_objects.UninitializedCoord) {
		if !(x == room.BotRight.X+1 && (room.BotRight.Y < y && y < room.TopLeft.Y)) {
			sl.ReportError(room.Doors, "door", "room.Doors[Right]", "door", "")
			return
		}
	}

	x, y = room.Doors[value_objects.Top].X, room.Doors[value_objects.Top].Y
	if !(x == value_objects.UninitializedCoord && y == value_objects.UninitializedCoord) {
		if !(y == room.TopLeft.Y+1 && (room.TopLeft.X < x && x < room.BotRight.X)) {
			sl.ReportError(room.Doors, "door", "room.Doors[Top]", "door", "")
			return
		}
	}

	x, y = room.Doors[value_objects.Bottom].X, room.Doors[value_objects.Bottom].Y
	if !(x == value_objects.UninitializedCoord && y == value_objects.UninitializedCoord) {
		if !(y == room.BotRight.Y-1 && (room.TopLeft.X < x && x < room.BotRight.X)) {
			sl.ReportError(room.Doors, "door", "room.Doors[Bottom]", "door", "")
			return
		}
	}

}
