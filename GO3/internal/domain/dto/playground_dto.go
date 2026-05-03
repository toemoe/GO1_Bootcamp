package dto

type PlaygroundDTO struct {
	UUID         string                `json:"uuid"`
	DungeonDTO   DungeonDTO            `json:"dungeon"`
	CharacterDTO CharacterDTO          `json:"character"`
	BackpackDTO  BackpackDTO           `json:"backpack"`
	MonstersDTO  []MonsterDTO          `json:"monsters"`
	FoodsDTO     []FoodPositionDTO     `json:"foods"`
	ScrollsDTO   []ScrollPositionDTO   `json:"scrolls"`
	ElixirsDTO   []ElixirPositionDTO   `json:"elixirs"`
	WeaponsDTO   []WeaponPositionDTO   `json:"weapons"`
	TreasuresDTO []TreasurePositionDTO `json:"treasure"`
}
