package entities

import (
	"s21_rogue/internal/domain/value_objects"
	"testing"

	"github.com/stretchr/testify/require"
)

// *******
func TestCorridorDirectHorizontal(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[3] = *value_objects.NewPosition(10, 5)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 6, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].X)
		require.Equal(t, 5, coordinates[i].Y)
	}
}

// *
// *
// *
// *
// *
func TestCorridorDirectVertical(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[3] = *value_objects.NewPosition(5, 10)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 6, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].Y)
		require.Equal(t, 5, coordinates[i].X)
	}
}

/*
***
..*
..***
*/
func TestCorridorSnakeHorizontal(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 5)
	corridor.Positions[2] = *value_objects.NewPosition(10, 3)
	corridor.Positions[3] = *value_objects.NewPosition(15, 3)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 13, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].X)
		require.Equal(t, 5, coordinates[i].Y)
	}
	require.Equal(t, 10, coordinates[6].X)
	require.Equal(t, 4, coordinates[6].Y)

	for i := 7; i < len(coordinates); i++ {
		require.Equal(t, i+3, coordinates[i].X)
		require.Equal(t, 3, coordinates[i].Y)
	}

}

/*
***
..*
..*
..*
..***
*/
func TestCorridorBigSnakeHorizontal(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(7, 8)
	corridor.Positions[1] = *value_objects.NewPosition(10, 8)
	corridor.Positions[2] = *value_objects.NewPosition(10, 3)
	corridor.Positions[3] = *value_objects.NewPosition(13, 3)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 12, len(coordinates))
	require.Equal(t, 7, coordinates[4].Y)
	require.Equal(t, 6, coordinates[5].Y)
	require.Equal(t, 4, coordinates[7].Y)
	require.Equal(t, 3, coordinates[8].Y)
}

/*
***
..***
*/
func TestCorridorSnakeHorizontalWithoutCenter(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 5)
	corridor.Positions[2] = *value_objects.NewPosition(10, 4)
	corridor.Positions[3] = *value_objects.NewPosition(15, 4)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 12, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].X)
		require.Equal(t, 5, coordinates[i].Y)
	}

	for i := 6; i < len(coordinates); i++ {
		require.Equal(t, i+4, coordinates[i].X)
		require.Equal(t, 4, coordinates[i].Y)
	}

}

/*
..***
..*
***
*/
func TestCorridorSnakeHorizontal2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 3)
	corridor.Positions[1] = *value_objects.NewPosition(10, 3)
	corridor.Positions[2] = *value_objects.NewPosition(10, 5)
	corridor.Positions[3] = *value_objects.NewPosition(15, 5)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 13, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].X)
		require.Equal(t, 3, coordinates[i].Y)
	}
	require.Equal(t, 10, coordinates[6].X)
	require.Equal(t, 4, coordinates[6].Y)

	for i := 7; i < len(coordinates); i++ {
		require.Equal(t, i+3, coordinates[i].X)
		require.Equal(t, 5, coordinates[i].Y)
	}

}

/*
..***
..*
..*
..*
***
*/
func TestCorridorBigSnakeHorizontal2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(7, 3)
	corridor.Positions[1] = *value_objects.NewPosition(10, 3)
	corridor.Positions[2] = *value_objects.NewPosition(10, 8)
	corridor.Positions[3] = *value_objects.NewPosition(13, 8)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 12, len(coordinates))
	require.Equal(t, 4, coordinates[4].Y)
	require.Equal(t, 5, coordinates[5].Y)
	require.Equal(t, 7, coordinates[7].Y)
	require.Equal(t, 8, coordinates[8].Y)
}

/*
..***
***
*/
func TestCorridorSnakeHorizontalWithoutCenter2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 4)
	corridor.Positions[1] = *value_objects.NewPosition(10, 4)
	corridor.Positions[2] = *value_objects.NewPosition(10, 5)
	corridor.Positions[3] = *value_objects.NewPosition(15, 5)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 12, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].X)
		require.Equal(t, 4, coordinates[i].Y)
	}

	for i := 6; i < len(coordinates); i++ {
		require.Equal(t, i+4, coordinates[i].X)
		require.Equal(t, 5, coordinates[i].Y)
	}

}

/*
*
*
***
..*
..*
*/
func TestCorridorSnakeVertical(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(10, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 10)
	corridor.Positions[2] = *value_objects.NewPosition(8, 10)
	corridor.Positions[3] = *value_objects.NewPosition(8, 15)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 13, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].Y)
		require.Equal(t, 10, coordinates[i].X)
	}
	require.Equal(t, 10, coordinates[6].Y)
	require.Equal(t, 9, coordinates[6].X)

	for i := 7; i < len(coordinates); i++ {
		require.Equal(t, i+3, coordinates[i].Y)
		require.Equal(t, 8, coordinates[i].X)
	}

}

/*
*
*
****
...*
...*
*/
func TestCorridorBigSnakeVertical(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(10, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 10)
	corridor.Positions[2] = *value_objects.NewPosition(6, 10)
	corridor.Positions[3] = *value_objects.NewPosition(6, 15)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 15, len(coordinates))

	require.Equal(t, 9, coordinates[6].X)
	require.Equal(t, 8, coordinates[7].X)
	require.Equal(t, 7, coordinates[8].X)
}

/*
..*
..*
***
*
*
*/
func TestCorridorSnakeVertical2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(8, 5)
	corridor.Positions[1] = *value_objects.NewPosition(8, 10)
	corridor.Positions[2] = *value_objects.NewPosition(10, 10)
	corridor.Positions[3] = *value_objects.NewPosition(10, 15)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 13, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].Y)
		require.Equal(t, 8, coordinates[i].X)
	}
	require.Equal(t, 10, coordinates[6].Y)
	require.Equal(t, 9, coordinates[6].X)

	for i := 7; i < len(coordinates); i++ {
		require.Equal(t, i+3, coordinates[i].Y)
		require.Equal(t, 10, coordinates[i].X)
	}

}

/*
...*
...*
****
*
*
*/
func TestCorridorBigSnakeVertical2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(6, 5)
	corridor.Positions[1] = *value_objects.NewPosition(6, 10)
	corridor.Positions[2] = *value_objects.NewPosition(10, 10)
	corridor.Positions[3] = *value_objects.NewPosition(10, 15)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 15, len(coordinates))

	require.Equal(t, 7, coordinates[6].X)
	require.Equal(t, 8, coordinates[7].X)
	require.Equal(t, 9, coordinates[8].X)
}

/*
.*
.*
**
*
*
*/
func TestCorridorSnakeVerticalWithoutCenter2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(9, 5)
	corridor.Positions[1] = *value_objects.NewPosition(9, 10)
	corridor.Positions[2] = *value_objects.NewPosition(10, 10)
	corridor.Positions[3] = *value_objects.NewPosition(10, 15)

	coordinates := corridor.GetAllCoordinates()
	require.Equal(t, 12, len(coordinates))
	for i := range 6 {
		require.Equal(t, i+5, coordinates[i].Y)
		require.Equal(t, 9, coordinates[i].X)
	}

	for i := 6; i < len(coordinates); i++ {
		require.Equal(t, i+4, coordinates[i].Y)
		require.Equal(t, 10, coordinates[i].X)
	}

}

// *******
func TestInCorridorDirectHorizontal(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[3] = *value_objects.NewPosition(10, 5)

	require.False(t, corridor.InCorridor(4, 5))
	require.False(t, corridor.InCorridor(5, 6))
	require.False(t, corridor.InCorridor(8, 4))
	require.False(t, corridor.InCorridor(11, 5))
	require.False(t, corridor.InCorridor(11, 4))

	require.True(t, corridor.InCorridor(5, 5))
	require.True(t, corridor.InCorridor(8, 5))
	require.True(t, corridor.InCorridor(10, 5))
}

// *
// *
// *
// *
// *
func TestInCorridorDirectVertical(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[3] = *value_objects.NewPosition(5, 10)

	require.False(t, corridor.InCorridor(5, 4))
	require.False(t, corridor.InCorridor(6, 5))
	require.False(t, corridor.InCorridor(4, 8))
	require.False(t, corridor.InCorridor(5, 11))
	require.False(t, corridor.InCorridor(4, 11))

	require.True(t, corridor.InCorridor(5, 5))
	require.True(t, corridor.InCorridor(5, 8))
	require.True(t, corridor.InCorridor(5, 10))
}

/*
***
..*
..***
*/
func TestInCorridorSnakeHorizontal(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 5)
	corridor.Positions[2] = *value_objects.NewPosition(10, 3)
	corridor.Positions[3] = *value_objects.NewPosition(15, 3)

	require.False(t, corridor.InCorridor(4, 5))
	require.False(t, corridor.InCorridor(5, 6))
	require.False(t, corridor.InCorridor(8, 4))
	require.False(t, corridor.InCorridor(11, 5))
	require.False(t, corridor.InCorridor(11, 4))
	require.False(t, corridor.InCorridor(9, 4))

	require.False(t, corridor.InCorridor(16, 3))
	require.False(t, corridor.InCorridor(16, 4))
	require.False(t, corridor.InCorridor(9, 3))

	require.True(t, corridor.InCorridor(5, 5))
	require.True(t, corridor.InCorridor(8, 5))
	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(10, 4))
	require.True(t, corridor.InCorridor(10, 3))
	require.True(t, corridor.InCorridor(12, 3))
	require.True(t, corridor.InCorridor(15, 3))
}

/*
***
..***
*/
func TestInCorridorSnakeHorizontalWithoutCenter(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 5)
	corridor.Positions[2] = *value_objects.NewPosition(10, 4)
	corridor.Positions[3] = *value_objects.NewPosition(15, 4)

	require.False(t, corridor.InCorridor(4, 5))
	require.False(t, corridor.InCorridor(5, 6))
	require.False(t, corridor.InCorridor(8, 4))
	require.False(t, corridor.InCorridor(9, 4))

	require.False(t, corridor.InCorridor(16, 3))
	require.False(t, corridor.InCorridor(16, 4))
	require.False(t, corridor.InCorridor(9, 3))

	require.True(t, corridor.InCorridor(5, 5))
	require.True(t, corridor.InCorridor(8, 5))
	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(10, 4))
	require.True(t, corridor.InCorridor(12, 4))
	require.True(t, corridor.InCorridor(15, 4))
}

/*
..***
..*
***
*/
func TestInCorridorSnakeHorizontal2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 3)
	corridor.Positions[1] = *value_objects.NewPosition(10, 3)
	corridor.Positions[2] = *value_objects.NewPosition(10, 5)
	corridor.Positions[3] = *value_objects.NewPosition(15, 5)

	require.False(t, corridor.InCorridor(4, 3))
	require.False(t, corridor.InCorridor(5, 4))
	require.False(t, corridor.InCorridor(8, 2))
	require.False(t, corridor.InCorridor(11, 3))
	require.False(t, corridor.InCorridor(11, 4))
	require.False(t, corridor.InCorridor(9, 5))

	require.False(t, corridor.InCorridor(16, 5))
	require.False(t, corridor.InCorridor(16, 6))
	require.False(t, corridor.InCorridor(9, 5))

	require.True(t, corridor.InCorridor(5, 3))
	require.True(t, corridor.InCorridor(8, 3))
	require.True(t, corridor.InCorridor(10, 3))
	require.True(t, corridor.InCorridor(10, 4))
	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(12, 5))
	require.True(t, corridor.InCorridor(15, 5))
}

/*
..***
***
*/
func TestInCorridorSnakeHorizontalWithoutCenter2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(5, 4)
	corridor.Positions[1] = *value_objects.NewPosition(10, 4)
	corridor.Positions[2] = *value_objects.NewPosition(10, 5)
	corridor.Positions[3] = *value_objects.NewPosition(15, 5)

	require.False(t, corridor.InCorridor(4, 4))
	require.False(t, corridor.InCorridor(5, 5))
	require.False(t, corridor.InCorridor(8, 3))
	require.False(t, corridor.InCorridor(9, 3))

	require.False(t, corridor.InCorridor(16, 5))
	require.False(t, corridor.InCorridor(16, 6))
	require.False(t, corridor.InCorridor(9, 5))

	require.True(t, corridor.InCorridor(5, 4))
	require.True(t, corridor.InCorridor(8, 4))
	require.True(t, corridor.InCorridor(10, 4))
	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(12, 5))
	require.True(t, corridor.InCorridor(15, 5))
}

/*
*
*
***
..*
..*
*/
func TestInCorridorSnakeVertical(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(10, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 10)
	corridor.Positions[2] = *value_objects.NewPosition(8, 10)
	corridor.Positions[3] = *value_objects.NewPosition(8, 15)

	require.False(t, corridor.InCorridor(10, 4))
	require.False(t, corridor.InCorridor(11, 5))
	require.False(t, corridor.InCorridor(9, 8))
	require.False(t, corridor.InCorridor(9, 11))

	require.False(t, corridor.InCorridor(8, 16))
	require.False(t, corridor.InCorridor(8, 9))

	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(10, 8))
	require.True(t, corridor.InCorridor(10, 10))
	require.True(t, corridor.InCorridor(9, 10))
	require.True(t, corridor.InCorridor(8, 10))
	require.True(t, corridor.InCorridor(8, 12))
	require.True(t, corridor.InCorridor(8, 15))
}

/*
*
*
**
.*
.*
*/
func TestInCorridorSnakeVerticalWithoutCenter(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(10, 5)
	corridor.Positions[1] = *value_objects.NewPosition(10, 10)
	corridor.Positions[2] = *value_objects.NewPosition(9, 10)
	corridor.Positions[3] = *value_objects.NewPosition(0, 15)

	require.False(t, corridor.InCorridor(10, 4))
	require.False(t, corridor.InCorridor(11, 5))
	require.False(t, corridor.InCorridor(9, 8))

	require.False(t, corridor.InCorridor(9, 16))
	require.False(t, corridor.InCorridor(9, 9))

	require.True(t, corridor.InCorridor(10, 5))
	require.True(t, corridor.InCorridor(10, 8))
	require.True(t, corridor.InCorridor(10, 10))
	require.True(t, corridor.InCorridor(9, 10))
	require.True(t, corridor.InCorridor(9, 12))
	require.True(t, corridor.InCorridor(9, 15))
}

/*
..*
..*
***
*
*
*/
func TestInCorridorSnakeVertical2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(8, 5)
	corridor.Positions[1] = *value_objects.NewPosition(8, 10)
	corridor.Positions[2] = *value_objects.NewPosition(10, 10)
	corridor.Positions[3] = *value_objects.NewPosition(10, 15)

	require.False(t, corridor.InCorridor(8, 4))
	require.False(t, corridor.InCorridor(9, 5))
	require.False(t, corridor.InCorridor(7, 8))
	require.False(t, corridor.InCorridor(7, 11))

	require.False(t, corridor.InCorridor(10, 16))
	require.False(t, corridor.InCorridor(10, 9))

	require.True(t, corridor.InCorridor(8, 5))
	require.True(t, corridor.InCorridor(8, 8))
	require.True(t, corridor.InCorridor(8, 10))
	require.True(t, corridor.InCorridor(9, 10))
	require.True(t, corridor.InCorridor(10, 10))
	require.True(t, corridor.InCorridor(10, 12))
	require.True(t, corridor.InCorridor(10, 15))
}

/*
.*
.*
**
*
*
*/
func TestInCorridorSnakeVerticalWithoutCenter2(t *testing.T) {
	corridor := NewCorridor()
	corridor.Positions[0] = *value_objects.NewPosition(8, 5)
	corridor.Positions[1] = *value_objects.NewPosition(8, 10)
	corridor.Positions[2] = *value_objects.NewPosition(9, 10)
	corridor.Positions[3] = *value_objects.NewPosition(9, 15)

	require.False(t, corridor.InCorridor(8, 4))
	require.False(t, corridor.InCorridor(9, 5))
	require.False(t, corridor.InCorridor(7, 8))
	require.False(t, corridor.InCorridor(7, 11))

	require.False(t, corridor.InCorridor(10, 16))
	require.False(t, corridor.InCorridor(10, 9))

	require.True(t, corridor.InCorridor(8, 5))
	require.True(t, corridor.InCorridor(8, 8))
	require.True(t, corridor.InCorridor(8, 10))
	require.True(t, corridor.InCorridor(9, 10))
	require.True(t, corridor.InCorridor(9, 12))
	require.True(t, corridor.InCorridor(9, 15))
}
