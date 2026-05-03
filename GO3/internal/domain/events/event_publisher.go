package events

type EventPublisher interface {
	Subscribe(handler EventHandler, types ...EventType)
	UnsubscribeAll()
	UnsubscribeFromAll(handler EventHandler)
	Notify(event Event)
}

type eventPublisherImpl struct {
	handlers map[EventType][]EventHandler
}

func NewEventPublisher() EventPublisher {
	return &eventPublisherImpl{handlers: make(map[EventType][]EventHandler)}
}

func (e *eventPublisherImpl) Subscribe(handler EventHandler, types ...EventType) {
	for _, t := range types {
		handlers, ok := e.handlers[t]
		if !ok {
			handlers = make([]EventHandler, 0)
		}
		handlers = append(handlers, handler)
		e.handlers[t] = handlers
	}
}
func (e *eventPublisherImpl) UnsubscribeAll() {
	e.handlers = make(map[EventType][]EventHandler)
}

func (e *eventPublisherImpl) UnsubscribeFromAll(handler EventHandler) {
	for typ, handlers := range e.handlers {
		newHandlers := []EventHandler{}
		for _, h := range handlers {
			if h != handler {
				newHandlers = append(newHandlers, h)
			}
		}

		e.handlers[typ] = newHandlers
	}
}

func (e *eventPublisherImpl) Notify(event Event) {
	if event.GetType() != NothingEvent {
		for _, handler := range e.handlers[event.GetType()] {
			handler.Handle(event)
		}
	}
}
