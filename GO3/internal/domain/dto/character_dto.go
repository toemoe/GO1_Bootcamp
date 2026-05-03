package dto

import "s21_rogue/internal/domain/value_objects"

type CharacterDTO struct {
	Position      value_objects.Position `json:"rooms"`
	CurrentHealth int                    `json:"current_health"`
	MaxHealth     int                    `json:"max_health"`
	Agile         int                    `json:"agile"`
	Strength      int                    `json:"strength"`
	Boosts        []BoostValue           `json:"boosts"`
}

const (
	AgileBoostType     string = "agile"
	MaxHealthBoostType string = "max_health"
	StrengthBoostType  string = "strength"
)

type BoostValue struct {
	BoostType  string `json:"boost_type"`
	BoostValue int    `json:"boost_value"`
	CountSteps int    `json:"count_steps"`
}
