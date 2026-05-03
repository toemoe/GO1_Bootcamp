package dto

type BackpackDTO struct {
	FoodsDTO     []FoodDTO     `json:"foods"`
	ScrollsDTO   []ScrollDTO   `json:"scrolls"`
	ElixirsDTO   []ElixirDTO   `json:"elixirs"`
	WeaponsDTO   []WeaponDTO   `json:"weapons"`
	TreasuresDTO []TreasureDTO `json:"treasures"`
}
