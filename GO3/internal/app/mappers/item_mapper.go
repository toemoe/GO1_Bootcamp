package mappers

import (
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
)

func MapToFoodDTO(food *entities.Food) *dto.FoodDTO {
	dto := dto.FoodDTO{}
	dto.UUID = food.GetUUID().String()
	dto.Label = food.Label
	dto.Health = food.Health

	return &dto
}

func MapFromFoodDTO(foodDTO *dto.FoodDTO) (*entities.Food, error) {
	return entities.NewFood(foodDTO.UUID, foodDTO.Label, foodDTO.Health)
}

func MapToScrollDTO(scroll *entities.Scroll) *dto.ScrollDTO {
	dto := dto.ScrollDTO{}
	dto.UUID = scroll.GetUUID().String()
	dto.Label = scroll.Label
	switch scroll.BoostType {
	case entities.AgileBoost:
		dto.BoostType = "agile"
	case entities.StrengthBoost:
		dto.BoostType = "strength"
	case entities.MaxHealthBoost:
		dto.BoostType = "max_health"
	}
	dto.Value = scroll.Value

	return &dto
}

func MapFromScrollDTO(scrollDTO *dto.ScrollDTO) (*entities.Scroll, error) {
	return entities.NewScroll(scrollDTO.UUID, scrollDTO.Label, scrollDTO.BoostType, scrollDTO.Value)
}

func MapToElixirDTO(elixir *entities.Elixir) *dto.ElixirDTO {
	dto := dto.ElixirDTO{}
	dto.UUID = elixir.GetUUID().String()
	dto.Label = elixir.Label
	switch elixir.BoostType {
	case entities.AgileBoost:
		dto.BoostType = "agile"
	case entities.StrengthBoost:
		dto.BoostType = "strength"
	case entities.MaxHealthBoost:
		dto.BoostType = "max_health"
	}
	dto.Value = elixir.Value
	dto.CountSteps = elixir.CountSteps

	return &dto
}

func MapFromElixirDTO(elixirDTO *dto.ElixirDTO) (*entities.Elixir, error) {
	return entities.NewElixir(elixirDTO.UUID, elixirDTO.Label, elixirDTO.BoostType, elixirDTO.Value, elixirDTO.CountSteps)
}

func MapToWeaponDTO(weapon *entities.Weapon) *dto.WeaponDTO {
	dto := dto.WeaponDTO{}
	dto.UUID = weapon.GetUUID().String()
	dto.Label = weapon.Label
	dto.Selected = weapon.IsSelected()
	dto.Strength = weapon.Strength

	return &dto
}

func MapFromWeaponDTO(weaponDTO *dto.WeaponDTO) (*entities.Weapon, error) {
	return entities.NewWeapon(weaponDTO.UUID, weaponDTO.Label, weaponDTO.Strength, weaponDTO.Selected)
}

func MapToTreasureDTO(food *entities.Treasure) *dto.TreasureDTO {
	dto := dto.TreasureDTO{}
	dto.UUID = food.GetUUID().String()
	dto.Label = food.Label
	dto.Score = food.Score

	return &dto
}

func MapFromTreasureDTO(treasureDTO *dto.TreasureDTO) (*entities.Treasure, error) {
	return entities.NewTreasure(treasureDTO.UUID, treasureDTO.Label, treasureDTO.Score)
}
