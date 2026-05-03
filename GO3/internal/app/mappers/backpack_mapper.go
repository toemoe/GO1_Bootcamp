package mappers

import (
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
)

func MapToBackpackDTO(backpack *entities.Backpack) *dto.BackpackDTO {
	dto := dto.BackpackDTO{}

	for e := backpack.Foods.Front(); e != nil; e = e.Next() {
		item := e.Value.(entities.Food)
		itemDTO := MapToFoodDTO(&item)
		dto.FoodsDTO = append(dto.FoodsDTO, *itemDTO)
	}

	for e := backpack.Scrolls.Front(); e != nil; e = e.Next() {
		item := e.Value.(entities.Scroll)
		itemDTO := MapToScrollDTO(&item)
		dto.ScrollsDTO = append(dto.ScrollsDTO, *itemDTO)
	}

	for e := backpack.Elixirs.Front(); e != nil; e = e.Next() {
		item := e.Value.(entities.Elixir)
		itemDTO := MapToElixirDTO(&item)
		dto.ElixirsDTO = append(dto.ElixirsDTO, *itemDTO)
	}

	for e := backpack.Weapons.Front(); e != nil; e = e.Next() {
		item := e.Value.(*entities.Weapon)
		itemDTO := MapToWeaponDTO(item)
		dto.WeaponsDTO = append(dto.WeaponsDTO, *itemDTO)
	}

	for e := backpack.Treasures.Front(); e != nil; e = e.Next() {
		item := e.Value.(entities.Treasure)
		itemDTO := MapToTreasureDTO(&item)
		dto.TreasuresDTO = append(dto.TreasuresDTO, *itemDTO)
	}

	return &dto
}

func MapFromBackpackDTO(BackpackDTO *dto.BackpackDTO) (*entities.Backpack, error) {
	backpack := entities.NewBackpack()

	for _, itemDTO := range BackpackDTO.FoodsDTO {
		item, err := MapFromFoodDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		backpack.Foods.PushBack(*item)
	}

	for _, itemDTO := range BackpackDTO.ScrollsDTO {
		item, err := MapFromScrollDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		backpack.Scrolls.PushBack(*item)
	}

	for _, itemDTO := range BackpackDTO.ElixirsDTO {
		item, err := MapFromElixirDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		backpack.Elixirs.PushBack(*item)
	}

	for _, itemDTO := range BackpackDTO.WeaponsDTO {
		item, err := MapFromWeaponDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		backpack.Weapons.PushBack(item)
	}

	for _, itemDTO := range BackpackDTO.TreasuresDTO {
		item, err := MapFromTreasureDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		backpack.Treasures.PushBack(*item)
	}

	return backpack, nil
}
