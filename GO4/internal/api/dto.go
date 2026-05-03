package api

import "github.com/google/uuid"

type GameDTO struct {
	ID    uuid.UUID `json:"id"`
	Board [3][3]int `json:"board"`
}
