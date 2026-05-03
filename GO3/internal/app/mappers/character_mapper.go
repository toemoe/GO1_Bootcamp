package mappers

import (
	"fmt"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
)

func MapToCharacterDTO(character *entities.Character) *dto.CharacterDTO {
	chDTO := dto.CharacterDTO{}
	chDTO.Position = *character.Position
	chDTO.CurrentHealth = character.GetHealth()
	chDTO.MaxHealth = character.GetMaxHealth()
	chDTO.Agile = character.GetAgile()
	chDTO.Strength = character.GetStrength()

	for e := character.GetAgileBoosts().Front(); e != nil; e = e.Next() {
		boost := e.Value.(entities.BoostTimes)
		chDTO.Boosts = append(chDTO.Boosts,
			dto.BoostValue{BoostType: dto.AgileBoostType,
				BoostValue: boost.BoostValue,
				CountSteps: boost.CountSteps})
	}

	for e := character.GetStrengthBoosts().Front(); e != nil; e = e.Next() {
		boost := e.Value.(entities.BoostTimes)
		chDTO.Boosts = append(chDTO.Boosts,
			dto.BoostValue{BoostType: dto.StrengthBoostType,
				BoostValue: boost.BoostValue,
				CountSteps: boost.CountSteps})
	}

	for e := character.GetMaxHealthBoosts().Front(); e != nil; e = e.Next() {
		boost := e.Value.(entities.BoostTimes)
		chDTO.Boosts = append(chDTO.Boosts,
			dto.BoostValue{BoostType: dto.MaxHealthBoostType,
				BoostValue: boost.BoostValue,
				CountSteps: boost.CountSteps})
	}

	return &chDTO
}

// Выбранное оружие в playground
func MapFromCharacterDTO(characterDTO *dto.CharacterDTO) (*entities.Character, error) {
	ch := entities.NewCharacter(&characterDTO.Position,
		characterDTO.CurrentHealth,
		characterDTO.MaxHealth,
		characterDTO.Agile,
		characterDTO.Strength)

	for _, v := range characterDTO.Boosts {
		switch v.BoostType {
		case dto.AgileBoostType:
			ch.AppendAgile(-v.BoostValue)
			ch.BoostAgile(v.BoostValue, v.CountSteps)
		case dto.StrengthBoostType:
			ch.AppendStrength(-v.BoostValue)
			ch.BoostStrength(v.BoostValue, v.CountSteps)
		case dto.MaxHealthBoostType:
			ch.AppendMaxHealth(-v.BoostValue)
			ch.BoostMaxHealth(v.BoostValue, v.CountSteps)
		default:
			return nil, fmt.Errorf("Incorrect type for boost value")
		}
	}

	return ch, nil
}
