package entities

import (
	"s21_rogue/internal/domain/commands"
	"s21_rogue/internal/domain/value_objects"
)

type Monster interface {
	GetAttackCommand() commands.AttackCommand
	TakeDamage(cmd commands.AttackCommand)
	GetMonsterType() MonsterType
	GetPosition() *value_objects.Position
	SetPosition(pos *value_objects.Position)
	CanAttack(position *value_objects.Position) bool
	IsChasing() bool
	SetChasing(bool)
	GetHealth() int
	GetHostility() int
	Walk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position)
	IsVisible() bool
}

type MonsterType int

const (
	ZombieType MonsterType = iota
	VampireType
	GhostType
	OgreType
	SnakeType
)

type HostilityType int

const (
	HighHostilityType   HostilityType = 8
	MediumHostilityType HostilityType = 5
	LowHostilityType    HostilityType = 3
)

const (
	HighAgileMonster   = 60
	MediumAgileMonster = 40
	LowAgileMonster    = 20
)

const (
	HighHealthMonster   = 40
	MediumHealthMonster = 30
	LowHealthMonster    = 20
)

const (
	HighStrengthMonster   = 10
	MediumStrengthMonster = 6
	LowStrengthMonster    = 3
)

type MonsterCommon struct {
	MonsterType MonsterType
	Hostility   int
	DamageableEntitiyBehaviour
	WalkableEnemyBehaviour
}

func NewMonster(position *value_objects.Position, monsterType MonsterType) Monster {
	switch monsterType {
	case ZombieType:
		return &MonsterCommon{
			MonsterType:                monsterType,
			Hostility:                  int(MediumHostilityType),
			DamageableEntitiyBehaviour: NewSimpleDamageableEntity(HighHealthMonster, HighHealthMonster, LowAgileMonster, MediumStrengthMonster),
			WalkableEnemyBehaviour:     NewSimpleWalkableEnemy(position)}
	case VampireType:
		return &MonsterCommon{
			MonsterType:                monsterType,
			Hostility:                  int(HighHostilityType),
			DamageableEntitiyBehaviour: NewVampireDamageableEntity(HighHealthMonster, HighAgileMonster, MediumStrengthMonster),
			WalkableEnemyBehaviour:     NewSimpleWalkableEnemy(position)}
	case GhostType:
		return &MonsterCommon{
			MonsterType:                monsterType,
			Hostility:                  int(LowHostilityType),
			DamageableEntitiyBehaviour: NewSimpleDamageableEntity(LowHealthMonster, LowHealthMonster, HighAgileMonster, LowStrengthMonster),
			WalkableEnemyBehaviour:     NewGhostWalkableEnemy(position)}
	case OgreType:
		return &MonsterCommon{
			MonsterType:                monsterType,
			Hostility:                  int(MediumHostilityType),
			DamageableEntitiyBehaviour: NewOrgeDamageableEntity(HighHealthMonster, LowAgileMonster, HighStrengthMonster),
			WalkableEnemyBehaviour:     NewOrgeWalkableEnemy(position)}
	case SnakeType:
		return &MonsterCommon{
			MonsterType:                monsterType,
			Hostility:                  int(HighHostilityType),
			DamageableEntitiyBehaviour: NewSnakeDamageableEntity(MediumHealthMonster, HighAgileMonster, MediumStrengthMonster),
			WalkableEnemyBehaviour:     NewSnakeWalkableEnemy(position)}
	}
	panic("nonimplemented")
}

func (mc *MonsterCommon) GetMonsterType() MonsterType {
	return mc.MonsterType

}

func (mc *MonsterCommon) GetHostility() int {
	return mc.Hostility
}
