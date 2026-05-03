package storage

import (
	"errors"
	"sync"
	domain "tic-tac-toe/internal/domain/model"

	"github.com/google/uuid"
)

type Memory struct {
	storage sync.Map
}

func NewMemoryRepository() *Memory {
	return &Memory{}
}

func (m *Memory) SaveGame(game domain.Game) error {
	item := &gameItem{
		data: ToEntity(game),
	}
	m.storage.Store(game.ID, item)
	return nil
}

func (m *Memory) GetGame(id uuid.UUID) (domain.Game, error) {
	val, ok := m.storage.Load(id)
	if !ok {
		return domain.Game{}, errors.New("game not found")
	}

	item := val.(*gameItem)
	item.mu.RLock()
	defer item.mu.RUnlock()
	return ToDomain(item.data), nil
}

func (m *Memory) UpdateGame(id uuid.UUID, updateFunc func(game domain.Game) (domain.Game, error)) error {
	val, ok := m.storage.Load(id)
	if !ok {
		return errors.New("game not found")
	}
	item := val.(*gameItem)
	item.mu.Lock()
	defer item.mu.Unlock()

	current := ToDomain(item.data)
	updated, err := updateFunc(current)
	if err != nil {
		return err
	}

	item.data = ToEntity(updated)
	return nil
}
