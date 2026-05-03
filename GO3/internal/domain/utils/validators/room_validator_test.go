package validators

import (
	"errors"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRoomValidator(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.Nil(t, err)
}

func TestRoomValidatorOutOfMapWidth(t *testing.T) {
	topLeft := value_objects.NewPosition(constants.MapWidth-constants.MinRoomWidth-1, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MapWidth, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorOutOfMapHeight(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MapHeight)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, constants.MapHeight-constants.MinRoomHeight-1)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorRoomLessThanZero(t *testing.T) {
	topLeft := value_objects.NewPosition(-1, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorRoomLessThanZero2(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, -1)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorBigRoomWidth(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorBigRoomHeight(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorMinRoomWidth(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MinRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorMinRoomHeight(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MinRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorUninitializedCoordDoor(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.Nil(t, err)
}

func TestRoomValidatorIncorrectLeftDoor(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Left] = *value_objects.NewPosition(-1, constants.MaxRoomHeight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectLeftDoor2(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Left] = *value_objects.NewPosition(0, constants.MaxRoomHeight-2)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectRightDoor(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Right] = *value_objects.NewPosition(constants.MaxRoomWidth, constants.MaxRoomHeight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectRightDoor2(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Right] = *value_objects.NewPosition(constants.MaxRoomWidth-1, constants.MaxRoomHeight-1)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectTopDoor(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Top] = *value_objects.NewPosition(1, constants.MaxRoomHeight+1)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectTopDoor2(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Top] = *value_objects.NewPosition(0, constants.MaxRoomHeight)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectBottomDoor(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Bottom] = *value_objects.NewPosition(1, 0)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}

func TestRoomValidatorIncorrectBottomDoor2(t *testing.T) {
	topLeft := value_objects.NewPosition(0, constants.MaxRoomHeight-1)
	botRight := value_objects.NewPosition(constants.MaxRoomWidth-1, 0)

	room := entities.NewRoom(*topLeft, *botRight)
	room.Doors[value_objects.Bottom] = *value_objects.NewPosition(0, -1)

	validate := validator.New()
	validate.RegisterStructValidation(RoomStructLevelValidation, entities.Room{})

	err := validate.Struct(room)
	assert.NotNil(t, err)

	var validateErrs validator.ValidationErrors
	assert.True(t, errors.As(err, &validateErrs))
}
