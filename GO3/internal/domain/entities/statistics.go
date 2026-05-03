package entities

import "github.com/google/uuid"

type Statistics struct {
	Id    uuid.UUID
	Level int

	MonsterDead    int
	GivenDamage    int
	RecievedDamage int

	FoodEaten     int
	ElixirsDrunk  int
	ScrollsRead   int
	TreasureFound int

	SteppedCount int
}

func NewStatistics(id uuid.UUID) *Statistics {
	return &Statistics{
		Id:    id,
		Level: 1,

		MonsterDead:    0,
		GivenDamage:    0,
		RecievedDamage: 0,

		FoodEaten:     0,
		ElixirsDrunk:  0,
		ScrollsRead:   0,
		TreasureFound: 0,

		SteppedCount: 0}
}
