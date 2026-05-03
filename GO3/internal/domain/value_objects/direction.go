package value_objects

type Directions int

const (
	Top Directions = iota
	Right
	Bottom
	Left
)

func GetOppositeDirection(dir Directions) Directions {
	switch dir {
	case Top:
		return Bottom
	case Bottom:
		return Top
	case Right:
		return Left
	default:
		return Right
	}
}
