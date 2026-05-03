package entities

import "s21_rogue/internal/domain/value_objects"

type Corridor struct {
	// Позиции всегда будут слева направо, снизу верх. Если корридор прямой - 1 и 2 позиция будет с координатами -1
	Positions  [4]value_objects.Position `json:"position"`
	VisitState value_objects.VisitState  `json:"visited_state"`
}

func NewCorridor() *Corridor {
	corridor := Corridor{VisitState: value_objects.NotVisitedState}
	for i := range 4 {
		corridor.Positions[i] = *value_objects.NewDefaultPosition()
	}
	return &corridor
}

func (c *Corridor) InCorridor(x, y int) bool {
	if c.IsDirectCorridor() {
		if c.IsHorizontalCorridor() {
			return y == c.Positions[0].Y && c.Positions[0].X <= x && x <= c.Positions[3].X
		}
		return x == c.Positions[0].X && c.Positions[0].Y <= y && y <= c.Positions[3].Y
	}

	if c.IsHorizontalCorridor() {
		if (y == c.Positions[0].Y && c.Positions[0].X <= x && x <= c.Positions[1].X) ||
			(y == c.Positions[2].Y && c.Positions[2].X <= x && x <= c.Positions[3].X) {
			return true
		}

		// ***
		//	 ***
		if c.Positions[1].Y < c.Positions[2].Y {
			return (x == c.Positions[1].X && c.Positions[1].Y <= y && y <= c.Positions[2].Y)
		}

		//	 ***
		// ***
		return (x == c.Positions[1].X && c.Positions[2].Y <= y && y <= c.Positions[1].Y)
	}

	// Vertical
	if (x == c.Positions[0].X && c.Positions[0].Y <= y && y <= c.Positions[1].Y) ||
		(x == c.Positions[2].X && c.Positions[2].Y <= y && y <= c.Positions[3].Y) {
		return true
	}

	//  *
	//  *
	// **
	// *
	// *
	if c.Positions[1].X < c.Positions[2].X {
		return (y == c.Positions[1].Y && c.Positions[1].X <= x && x <= c.Positions[2].X)
	}

	// *
	// *
	// **
	//  *
	//  *
	return (y == c.Positions[1].Y && c.Positions[2].X <= x && x <= c.Positions[1].X)

}

func (c *Corridor) IsDirectCorridor() bool {
	return c.Positions[1].X == value_objects.UninitializedCoord
}

func (c *Corridor) IsHorizontalCorridor() bool {
	compareIndex := 1
	if c.IsDirectCorridor() {
		compareIndex = 3

	}
	return c.Positions[0].Y == c.Positions[compareIndex].Y
}

func (c *Corridor) GetAllCoordinates() []*value_objects.Position {
	res := make([]*value_objects.Position, 0)
	if c.IsDirectCorridor() && c.IsHorizontalCorridor() {
		for x := c.Positions[0].X; x <= c.Positions[3].X; x++ {
			res = append(res, value_objects.NewPosition(x, c.Positions[0].Y))
		}
		return res
	} else if c.IsDirectCorridor() && !c.IsHorizontalCorridor() {
		for y := c.Positions[0].Y; y <= c.Positions[3].Y; y++ {
			res = append(res, value_objects.NewPosition(c.Positions[0].X, y))
		}
		return res
	} else if !c.IsDirectCorridor() && c.IsHorizontalCorridor() {
		for x := c.Positions[0].X; x <= c.Positions[1].X; x++ {
			res = append(res, value_objects.NewPosition(x, c.Positions[0].Y))
		}
		if c.Positions[1].Y < c.Positions[2].Y {
			for y := c.Positions[1].Y + 1; y < c.Positions[2].Y; y++ {
				res = append(res, value_objects.NewPosition(c.Positions[1].X, y))
			}
		} else {
			for y := c.Positions[1].Y - 1; y > c.Positions[2].Y; y-- {
				res = append(res, value_objects.NewPosition(c.Positions[1].X, y))
			}
		}
		for x := c.Positions[2].X; x <= c.Positions[3].X; x++ {
			res = append(res, value_objects.NewPosition(x, c.Positions[2].Y))
		}
		return res
	} else {
		for y := c.Positions[0].Y; y <= c.Positions[1].Y; y++ {
			res = append(res, value_objects.NewPosition(c.Positions[0].X, y))
		}
		if c.Positions[1].X < c.Positions[2].X {
			for x := c.Positions[1].X + 1; x < c.Positions[2].X; x++ {
				res = append(res, value_objects.NewPosition(x, c.Positions[1].Y))
			}
		} else {
			for x := c.Positions[1].X - 1; x > c.Positions[2].X; x-- {
				res = append(res, value_objects.NewPosition(x, c.Positions[1].Y))
			}
		}
		for y := c.Positions[2].Y; y <= c.Positions[3].Y; y++ {
			res = append(res, value_objects.NewPosition(c.Positions[2].X, y))
		}
		return res
	}
}

func (c *Corridor) GetIndexByPosition(pos *value_objects.Position) int {
	allCoordinates := c.GetAllCoordinates()
	for i := range allCoordinates {
		if allCoordinates[i].IsEqual(pos) {
			return i
		}
	}
	return -1
}
