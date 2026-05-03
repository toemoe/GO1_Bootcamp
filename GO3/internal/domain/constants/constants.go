package constants

const (
	RoomsPerSide    = 3
	OffsetPerSector = 4

	MaxRoomWidth  = 28
	MinRoomWidth  = 6
	MaxRoomHeight = 8
	MinRoomHeight = 3

	MapWidth  = (RoomsPerSide * MaxRoomWidth) + (RoomsPerSide-1)*OffsetPerSector
	MapHeight = (RoomsPerSide * MaxRoomHeight) + (RoomsPerSide-1)*OffsetPerSector

	MaxCorridorCounts = (RoomsPerSide - 1) * RoomsPerSide * 2
	MaxDeletedRooms   = 4
	MinDeletedRooms   = 2

	MaxGameLevel   = 21
	RadiusFogOfWar = 5.0
)
