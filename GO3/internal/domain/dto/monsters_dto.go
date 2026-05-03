package dto

import "s21_rogue/internal/domain/value_objects"

type MonsterTypeDTO string

const (
	ZombieTypeDTO  MonsterTypeDTO = "zombie"
	VampireTypeDTO MonsterTypeDTO = "vampire"
	GhostTypeDTO   MonsterTypeDTO = "ghost"
	OgreTypeDTO    MonsterTypeDTO = "ogre"
	SnakeTypeDTO   MonsterTypeDTO = "snake"
)

type MonsterDTO struct {
	Position       value_objects.Position `json:"position"`
	MonsterTypeDTO string                 `json:"monster_type"`
}
