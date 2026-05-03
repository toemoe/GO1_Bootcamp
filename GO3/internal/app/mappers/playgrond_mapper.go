package mappers

import (
	"container/list"
	"s21_rogue/internal/domain/aggregates"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/value_objects"

	"github.com/google/uuid"
)

func MapToPlaygroundDTO(playground *aggregates.Playground) *dto.PlaygroundDTO {
	playgroundDTO := dto.PlaygroundDTO{}
	playgroundDTO.UUID = playground.Id.String()
	playgroundDTO.DungeonDTO = *MapToDungeonDTO(playground.Dungeon)
	playgroundDTO.CharacterDTO = *MapToCharacterDTO(playground.Character)
	playgroundDTO.BackpackDTO = *MapToBackpackDTO(playground.Backpack)
	playgroundDTO.MonstersDTO = make([]dto.MonsterDTO, 0)
	for e := playground.Monsters.Front(); e != nil; e = e.Next() {
		item := e.Value.(entities.Monster)
		itemDTO := MapToMonsterDTO(item)
		playgroundDTO.MonstersDTO = append(playgroundDTO.MonstersDTO, *itemDTO)
	}

	for k, v := range playground.Foods {
		itemDTO := MapToFoodDTO(&v)
		itemPosDTO := dto.FoodPositionDTO{FoodDTO: *itemDTO, Position: k}
		playgroundDTO.FoodsDTO = append(playgroundDTO.FoodsDTO, itemPosDTO)
	}

	for k, v := range playground.Scrolls {
		itemDTO := MapToScrollDTO(&v)
		itemPosDTO := dto.ScrollPositionDTO{ScrollDTO: *itemDTO, Position: k}
		playgroundDTO.ScrollsDTO = append(playgroundDTO.ScrollsDTO, itemPosDTO)
	}

	for k, v := range playground.Elixirs {
		itemDTO := MapToElixirDTO(&v)
		itemPosDTO := dto.ElixirPositionDTO{ElixirDTO: *itemDTO, Position: k}
		playgroundDTO.ElixirsDTO = append(playgroundDTO.ElixirsDTO, itemPosDTO)
	}

	for k, v := range playground.Weapons {
		itemDTO := MapToWeaponDTO(v)
		itemPosDTO := dto.WeaponPositionDTO{WeaponDTO: *itemDTO, Position: k}
		playgroundDTO.WeaponsDTO = append(playgroundDTO.WeaponsDTO, itemPosDTO)
	}

	for k, v := range playground.Treasures {
		itemDTO := MapToTreasureDTO(&v)
		itemPosDTO := dto.TreasurePositionDTO{TreasureDTO: *itemDTO, Position: k}
		playgroundDTO.TreasuresDTO = append(playgroundDTO.TreasuresDTO, itemPosDTO)
	}

	return &playgroundDTO
}

func MapFromPlaygroundDTO(playgroundDTO *dto.PlaygroundDTO) (*aggregates.Playground, error) {
	uuid, err := uuid.Parse(playgroundDTO.UUID)
	if err != nil {
		return nil, err
	}

	dungeon, err := MapFromDungeonDTO(&playgroundDTO.DungeonDTO)
	if err != nil {
		return nil, err
	}

	character, err := MapFromCharacterDTO(&playgroundDTO.CharacterDTO)
	if err != nil {
		return nil, err
	}

	backpack, err := MapFromBackpackDTO(&playgroundDTO.BackpackDTO)
	if err != nil {
		return nil, err
	}

	for e := backpack.Weapons.Front(); e != nil; e = e.Next() {
		w := e.Value.(*entities.Weapon)
		if w.IsSelected() {
			character.SetWeaponStrength(w.Strength)
		}
	}

	monsters := list.New()
	for _, itemDTO := range playgroundDTO.MonstersDTO {
		item, err := MapFromMonsterDTO(&itemDTO)
		if err != nil {
			return nil, err
		}
		monsters.PushBack(item)
	}

	foods := make(map[value_objects.Position]entities.Food)
	for _, itemDTO := range playgroundDTO.FoodsDTO {
		item, err := MapFromFoodDTO(&itemDTO.FoodDTO)
		if err != nil {
			return nil, err
		}
		foods[itemDTO.Position] = *item
	}

	scrolls := make(map[value_objects.Position]entities.Scroll)
	for _, itemDTO := range playgroundDTO.ScrollsDTO {
		item, err := MapFromScrollDTO(&itemDTO.ScrollDTO)
		if err != nil {
			return nil, err
		}
		scrolls[itemDTO.Position] = *item
	}

	elixirs := make(map[value_objects.Position]entities.Elixir)
	for _, itemDTO := range playgroundDTO.ElixirsDTO {
		item, err := MapFromElixirDTO(&itemDTO.ElixirDTO)
		if err != nil {
			return nil, err
		}
		elixirs[itemDTO.Position] = *item
	}

	weapons := make(map[value_objects.Position]*entities.Weapon)
	for _, itemDTO := range playgroundDTO.WeaponsDTO {
		item, err := MapFromWeaponDTO(&itemDTO.WeaponDTO)
		if err != nil {
			return nil, err
		}
		weapons[itemDTO.Position] = item
	}

	return aggregates.NewPlayground(uuid, dungeon, character, backpack, monsters, foods, scrolls, elixirs, weapons), nil
}
