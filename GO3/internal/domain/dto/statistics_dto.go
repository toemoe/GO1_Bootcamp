package dto

type StatisticsDTO struct {
	Id    string `json:"uuid"`
	Level int    `json:"level"`

	MonsterDead    int `json:"monster_dead"`
	GivenDamage    int `json:"given_damage"`
	RecievedDamage int `json:"recieved_damage"`

	FoodEaten     int `json:"food_eaten"`
	ElixirsDrunk  int `json:"elixirs_drunk"`
	ScrollsRead   int `json:"scrolls_read"`
	TreasureFound int `json:"treasure_found"`

	SteppedCount int `json:"stepped_count"`
}

type StatisticsSliceDTO struct {
	Statistics []StatisticsDTO `json:"statistics"`
}
