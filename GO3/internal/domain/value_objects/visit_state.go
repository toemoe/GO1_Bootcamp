package value_objects

type VisitState int

const (
	VisitedState VisitState = iota
	NotVisitedState
	CurrentSpace
)
