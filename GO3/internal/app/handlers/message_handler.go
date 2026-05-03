package handlers

import (
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/utils"
)

type massageHandler struct {
	log utils.Logger
}

func NewMessageHandler(log utils.Logger) events.EventHandler {
	return &massageHandler{log: log}
}

func (m *massageHandler) Handle(event events.Event) {
	switch event.GetType() {
	case events.FoodEatenEvent:
		m.log.GameMessage("Character ate food.")
	case events.ElixirsDrunkEvent:
		m.log.GameMessage("Character drank an elixir.")
	case events.ScrollsReadEvent:
		m.log.GameMessage("Character read a scroll.")
	case events.MonsterDeadEvent:
		m.log.GameMessage("Monster was defeated by the character.")
	case events.MonsterAttacked:
		m.log.GameMessage("Monster attacked the character.")
	case events.CharacterAttacked:
		m.log.GameMessage("Character attacked a monster.")
	case events.LevelChangedEvent:
		m.log.GameMessage("Level has changed.")
	}
}
