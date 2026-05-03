package value_objects

type Action int

const (
	StartAction Action = iota
	LoadAction
	TerminateAction
	BackAction
	StatisticsAction

	UpAction
	RightAction
	DownAction
	LeftAction

	FoodAction
	ScrollAction
	ElixirAction
	WeaponAction

	Nothing
)

func (a Action) IsMovingAction() bool {
	return UpAction <= a && a <= LeftAction
}
