package events

type EventType int

const (
	FoodEatenEvent EventType = iota
	ElixirsDrunkEvent
	ScrollsReadEvent
	TreasureFoundEvent

	NextLevelEvent
	MonsterDeadEvent
	LevelChangedEvent

	MonsterAttacked
	CharacterAttacked

	CharacterStepped

	NothingEvent
)

type Event interface {
	GetType() EventType
}

type eventImpl struct {
	EventType EventType
}

func (e eventImpl) GetType() EventType {
	return e.EventType
}

func NewEvent(e EventType) Event {
	return eventImpl{EventType: e}
}
