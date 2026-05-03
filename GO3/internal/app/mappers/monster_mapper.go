package mappers

import (
	"fmt"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
)

func MapToMonsterDTO(monster entities.Monster) *dto.MonsterDTO {
	monsterDTO := dto.MonsterDTO{}
	monsterDTO.Position = *monster.GetPosition()
	switch monster.GetMonsterType() {
	case entities.ZombieType:
		monsterDTO.MonsterTypeDTO = string(dto.ZombieTypeDTO)
	case entities.VampireType:
		monsterDTO.MonsterTypeDTO = string(dto.VampireTypeDTO)
	case entities.GhostType:
		monsterDTO.MonsterTypeDTO = string(dto.GhostTypeDTO)
	case entities.OgreType:
		monsterDTO.MonsterTypeDTO = string(dto.OgreTypeDTO)
	case entities.SnakeType:
		monsterDTO.MonsterTypeDTO = string(dto.SnakeTypeDTO)
	default:
		panic("incorrect monster type")
	}

	return &monsterDTO
}

func MapFromMonsterDTO(monsterDTO *dto.MonsterDTO) (entities.Monster, error) {
	var mt entities.MonsterType
	switch monsterDTO.MonsterTypeDTO {
	case string(dto.ZombieTypeDTO):
		mt = entities.ZombieType
	case string(dto.VampireTypeDTO):
		mt = entities.VampireType
	case string(dto.GhostTypeDTO):
		mt = entities.GhostType
	case string(dto.OgreTypeDTO):
		mt = entities.OgreType
	case string(dto.SnakeTypeDTO):
		mt = entities.SnakeType
	default:
		return nil, fmt.Errorf("Incorrect type")
	}

	return entities.NewMonster(&monsterDTO.Position, mt), nil
}
