package validators

import (
	"errors"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/utils/generators"
	"s21_rogue/internal/domain/value_objects"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestDungeonValidator(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.Nil(t, err)
}

func TestDungeonValidatorDive(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	dungeon.Rooms[0][0].BotRight.X = constants.MapWidth * 2

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeHorizontalDoor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	foundDoor := false
	for i := 0; i < constants.RoomsPerSide && !foundDoor; i++ {
		for j := 0; j < constants.RoomsPerSide && !foundDoor; j++ {
			rightDoor := dungeon.Rooms[i][j].Doors[value_objects.Right]
			if !rightDoor.IsDefault() {
				dungeon.Rooms[i][j+1].Doors[value_objects.Left].X = value_objects.UninitializedCoord
				dungeon.Rooms[i][j+1].Doors[value_objects.Left].Y = value_objects.UninitializedCoord
				foundDoor = true
			}
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeHorizontalDoor2(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	foundDoor := false
	for i := 0; i < constants.RoomsPerSide && !foundDoor; i++ {
		for j := 1; j < constants.RoomsPerSide && !foundDoor; j++ {
			leftDoor := dungeon.Rooms[i][j].Doors[value_objects.Left]
			if !leftDoor.IsDefault() {
				dungeon.Rooms[i][j-1].Doors[value_objects.Right].X = value_objects.UninitializedCoord
				dungeon.Rooms[i][j-1].Doors[value_objects.Right].Y = value_objects.UninitializedCoord
				foundDoor = true
			}
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeHorizontalDoorLeft(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	dungeon.Rooms[0][0].Doors[value_objects.Left].X = dungeon.Rooms[0][0].TopLeft.X - 1
	dungeon.Rooms[0][0].Doors[value_objects.Left].Y = dungeon.Rooms[0][0].TopLeft.Y - 1

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeHorizontalDoorRight(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	dungeon.Rooms[constants.RoomsPerSide-1][constants.RoomsPerSide-1].Doors[value_objects.Right].X = dungeon.Rooms[constants.RoomsPerSide-1][constants.RoomsPerSide-1].BotRight.X + 1
	dungeon.Rooms[constants.RoomsPerSide-1][constants.RoomsPerSide-1].Doors[value_objects.Right].Y = dungeon.Rooms[constants.RoomsPerSide-1][constants.RoomsPerSide-1].BotRight.Y + 1

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeVerticalDoorTop(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	dungeon.Rooms[constants.RoomsPerSide-1][0].Doors[value_objects.Top].X = dungeon.Rooms[constants.RoomsPerSide-1][0].TopLeft.X + 1
	dungeon.Rooms[constants.RoomsPerSide-1][0].Doors[value_objects.Top].Y = dungeon.Rooms[constants.RoomsPerSide-1][0].TopLeft.Y + 1

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeVerticalDoorBot(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	dungeon.Rooms[0][0].Doors[value_objects.Bottom].X = dungeon.Rooms[0][0].BotRight.X - 1
	dungeon.Rooms[0][0].Doors[value_objects.Bottom].Y = dungeon.Rooms[0][0].BotRight.Y - 1

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeVerticalDoor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	foundDoor := false
	for i := 0; i < constants.RoomsPerSide && !foundDoor; i++ {
		for j := 0; j < constants.RoomsPerSide && !foundDoor; j++ {
			topDoor := dungeon.Rooms[i][j].Doors[value_objects.Top]
			if !topDoor.IsDefault() {
				dungeon.Rooms[i+1][j].Doors[value_objects.Bottom].X = value_objects.UninitializedCoord
				dungeon.Rooms[i+1][j].Doors[value_objects.Bottom].Y = value_objects.UninitializedCoord
				foundDoor = true
			}
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonValidatorDelOppositeBottomDoor2(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	foundDoor := false
	for i := 1; i < constants.RoomsPerSide && !foundDoor; i++ {
		for j := 0; j < constants.RoomsPerSide && !foundDoor; j++ {
			botDoor := dungeon.Rooms[i][j].Doors[value_objects.Bottom]
			if !botDoor.IsDefault() {
				dungeon.Rooms[i-1][j].Doors[value_objects.Top].X = value_objects.UninitializedCoord
				dungeon.Rooms[i-1][j].Doors[value_objects.Top].Y = value_objects.UninitializedCoord
				foundDoor = true
			}
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonDuplicateHorizontalCorridor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	var foundHorizontalStartPos value_objects.Position
	isFound := false
	for i := 1; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if !dungeon.Rooms[i][j].Doors[value_objects.Right].IsDefault() {
				foundHorizontalStartPos.X = dungeon.Rooms[i][j].Doors[value_objects.Right].X
				foundHorizontalStartPos.Y = dungeon.Rooms[i][j].Doors[value_objects.Right].Y
				isFound = true
			}
		}
	}

	var addedCorridor entities.Corridor
	for _, c := range dungeon.Corridors {
		if c.Positions[0].X == foundHorizontalStartPos.X && c.Positions[0].Y == foundHorizontalStartPos.Y {
			addedCorridor = c
			break
		}
	}
	dungeon.Corridors = append(dungeon.Corridors, addedCorridor)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonDeletedHorizontalCorridor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	var foundHorizontalStartPos value_objects.Position
	isFound := false
	for i := 0; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if !dungeon.Rooms[i][j].Doors[value_objects.Right].IsDefault() {
				foundHorizontalStartPos.X = dungeon.Rooms[i][j].Doors[value_objects.Right].X
				foundHorizontalStartPos.Y = dungeon.Rooms[i][j].Doors[value_objects.Right].Y
				isFound = true
			}
		}
	}

	var deletedIndex int
	for i, c := range dungeon.Corridors {
		if c.Positions[0].X == foundHorizontalStartPos.X && c.Positions[0].Y == foundHorizontalStartPos.Y {
			deletedIndex = i
			break
		}
	}
	dungeon.Corridors = append(dungeon.Corridors[:deletedIndex], dungeon.Corridors[deletedIndex+1:]...)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonDeletedVerticalCorridor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	var foundVerticalStartPos value_objects.Position
	isFound := false
	for i := 0; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if !dungeon.Rooms[i][j].Doors[value_objects.Top].IsDefault() {
				foundVerticalStartPos.X = dungeon.Rooms[i][j].Doors[value_objects.Top].X
				foundVerticalStartPos.Y = dungeon.Rooms[i][j].Doors[value_objects.Top].Y
				isFound = true
			}
		}
	}

	var deletedIndex int
	for i, c := range dungeon.Corridors {
		if c.Positions[0].X == foundVerticalStartPos.X && c.Positions[0].Y == foundVerticalStartPos.Y {
			deletedIndex = i
			break
		}
	}
	dungeon.Corridors = append(dungeon.Corridors[:deletedIndex], dungeon.Corridors[deletedIndex+1:]...)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonChangedVerticalCorridor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	var foundVerticalStartPos value_objects.Position
	isFound := false
	for i := 0; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if !dungeon.Rooms[i][j].Doors[value_objects.Top].IsDefault() {
				foundVerticalStartPos.X = dungeon.Rooms[i][j].Doors[value_objects.Top].X
				foundVerticalStartPos.Y = dungeon.Rooms[i][j].Doors[value_objects.Top].Y
				isFound = true
			}
		}
	}

	for i, c := range dungeon.Corridors {
		if c.Positions[0].X == foundVerticalStartPos.X && c.Positions[0].Y == foundVerticalStartPos.Y {
			dungeon.Corridors[i].Positions[0].Y--
			break
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonChangedInnerVerticalCorridor(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	var foundVerticalStartPos value_objects.Position
	isFound := false
	for i := 0; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if !dungeon.Rooms[i][j].Doors[value_objects.Top].IsDefault() {
				foundVerticalStartPos.X = dungeon.Rooms[i][j].Doors[value_objects.Top].X
				foundVerticalStartPos.Y = dungeon.Rooms[i][j].Doors[value_objects.Top].Y
				isFound = true
			}
		}
	}

	for i, c := range dungeon.Corridors {
		if c.Positions[0].X == foundVerticalStartPos.X && c.Positions[0].Y == foundVerticalStartPos.Y {
			dungeon.Corridors[i].Positions[1].X--
			break
		}
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestDungeonIncorrectSector(t *testing.T) {
	logger := utils.NewNoopLogger()
	dg := generators.NewDungeonGenerator(logger)
	dungeon := dg.GenerateDungeon()

	temp := dungeon.Rooms[0][0].BotRight.X
	dungeon.Rooms[0][0].BotRight.X = constants.MaxRoomWidth + 1
	if !dungeon.Rooms[0][0].Doors[value_objects.Right].IsDefault() {
		for i, c := range dungeon.Corridors {
			if c.Positions[0].X == dungeon.Rooms[0][0].Doors[value_objects.Right].X &&
				c.Positions[0].Y == dungeon.Rooms[0][0].Doors[value_objects.Right].Y {
				dungeon.Corridors[i].Positions[0].X = constants.MaxRoomWidth + 2
				break
			}
		}
		dungeon.Rooms[0][0].Doors[value_objects.Right].X = constants.MaxRoomWidth + 2
	}
	offset := constants.MaxRoomWidth + 1 - temp
	dungeon.Rooms[0][0].TopLeft.X += offset
	if !dungeon.Rooms[0][0].Doors[value_objects.Top].IsDefault() {
		for i, c := range dungeon.Corridors {
			if c.Positions[0].X == dungeon.Rooms[0][0].Doors[value_objects.Top].X &&
				c.Positions[0].Y == dungeon.Rooms[0][0].Doors[value_objects.Top].Y {
				dungeon.Corridors[i].Positions[0].X += offset
				if !dungeon.Corridors[i].Positions[1].IsDefault() {
					dungeon.Corridors[i].Positions[1].X += offset
				}
				break
			}
		}
		dungeon.Rooms[0][0].Doors[value_objects.Top].X += offset
	}

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})
	validate.RegisterStructValidation(DungeonStructLevelValidation, entities.Dungeon{})

	err := validate.Struct(dungeon)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}
