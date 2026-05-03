package value_objects

// ^ y
// |
// |
// |-----> x

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

const (
	UninitializedCoord = -1
)

func NewPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}

func NewDefaultPosition() *Position {
	return &Position{X: UninitializedCoord, Y: UninitializedCoord}
}

func (p *Position) IsDefault() bool {
	return p.X == UninitializedCoord && p.Y == UninitializedCoord
}

func (p *Position) IsEqual(p2 *Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}
