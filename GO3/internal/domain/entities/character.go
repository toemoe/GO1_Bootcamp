package entities

import (
	"s21_rogue/internal/domain/value_objects"
)

type Character struct {
	Position *value_objects.Position
	DamageableEntitiyBehaviour
}

func NewCharacter(position *value_objects.Position, health, maxHealth, agile, strength int) *Character {
	return &Character{Position: position,
		DamageableEntitiyBehaviour: NewSimpleDamageableEntity(health, maxHealth, agile, strength)}
}

func (c *Character) GetNextActionPos(action value_objects.Action) *value_objects.Position {

	nextPos := value_objects.NewPosition(c.Position.X, c.Position.Y)
	if c.IsTranquilize() {
		c.SetTranquilize(false)
		return nextPos
	}
	switch action {
	case value_objects.UpAction:
		nextPos.Y++
	case value_objects.DownAction:
		nextPos.Y--
	case value_objects.RightAction:
		nextPos.X++
	case value_objects.LeftAction:
		nextPos.X--
	}
	return nextPos
}

func (c *Character) SetPosition(position *value_objects.Position) {
	c.Position = position
}

func (c *Character) IsDeath() bool {
	return c.GetHealth() <= 0
}
