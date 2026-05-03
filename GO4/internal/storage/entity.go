package storage

import (
	"sync"

	"github.com/google/uuid"
)

type GameEntity struct {
	ID    uuid.UUID
	Board [3][3]int
}

type gameItem struct {
	mu   sync.RWMutex
	data GameEntity
}
