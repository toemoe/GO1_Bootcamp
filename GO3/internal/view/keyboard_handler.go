package view

import (
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/value_objects"
)

type Handler interface {
	Next(handler Handler)
	Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool
}

type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) Next(handler Handler) {
	h.nextHandler = handler
}

func (h *BaseHandler) Handle(action value_objects.Action, gameInfo *dto.GameInfoDTO, playground *dto.PlaygroundDTO) bool {
	if h.nextHandler != nil {
		return h.nextHandler.Handle(action, gameInfo, playground)
	}
	return false
}
