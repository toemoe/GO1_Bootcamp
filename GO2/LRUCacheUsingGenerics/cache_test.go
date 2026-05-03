package cache

import (
	"testing"
)

func TestCacheEviction(t *testing.T) {
	cache := NewCache[string, int](2)

	cache.Set("a", 1)
	cache.Set("b", 2)

	cache.Get("a")
	cache.Set("c", 3)

	if _, ok := cache.Get("b"); ok {
		t.Fatal("expected 'b' to be evicted")
	}
}

func TestZeroCapacity(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when capacity <= 0")
		}
	}()
	_ = NewCache[int, int](0)
}

func TestCapacityOverflow(t *testing.T) {
	cache := NewCache[int, int](1)

	cache.Set(1, 1)
	cache.Set(2, 2)

	if _, ok := cache.Get(1); ok {
		t.Fatal("expected key 1 to be evicted")
	}
}
