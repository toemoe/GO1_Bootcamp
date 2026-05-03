package events

type EventHandler interface {
	Handle(event Event)
}
