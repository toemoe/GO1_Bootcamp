package cache

import (
	"container/list"
	"sync"
)

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]*list.Element
	list     *list.List
	mu       sync.RWMutex
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("Capacity must be greater than 0")
	}

	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
		list:     list.New(),
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, exists := c.items[key]
	if !exists {
		var zero V
		return zero, false
	}

	c.list.MoveToFront(element)
	return element.Value.(*entry[K, V]).value, true
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if element, exists := c.items[key]; exists {
		element.Value.(*entry[K, V]).value = value
		c.list.MoveToFront(element)
		return
	}

	if c.list.Len() >= c.capacity {
		c.evict()
	}

	element := c.list.PushFront(&entry[K, V]{
		key:   key,
		value: value,
	})

	c.items[key] = element
}

func (c *Cache[K, V]) evict() {
	last := c.list.Back()
	if last == nil {
		return
	}

	entry := last.Value.(*entry[K, V])

	delete(c.items, entry.key)
	c.list.Remove(last)

}

func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key := range c.items {
		delete(c.items, key)
	}

	c.list.Init()
}
