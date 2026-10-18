package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	pokedex    map[string]cacheEntry
	mu         *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	var c = Cache{
		cacheEntry: make(map[string]cacheEntry),
		pokedex:    make(map[string]cacheEntry),
		mu:         &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) error {
	if len(key) == 0 {
		return fmt.Errorf("key error")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cacheEntry[key]
	if !ok {
		return []byte{}, false
	}
	val := entry.val
	return val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	for range ticker.C {

		c.mu.Lock()

		now := time.Now()
		for key, entry := range c.cacheEntry {
			if now.Sub(entry.createdAt) > interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
