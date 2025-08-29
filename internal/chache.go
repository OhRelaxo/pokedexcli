package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mux   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (Cache, error) {
	var c Cache
	if interval <= 0 {
		return c, fmt.Errorf("interval cant be 0")
	}
	c.mux = &sync.RWMutex{}
	c.entry = make(map[string]cacheEntry)
	go c.reapLoop(interval)
	return c, nil
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entry[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	if entry, ok := c.entry[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) PrintAllKeys() {
	for k := range c.entry {
		fmt.Printf("Key: %v\n", k)
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	watch := time.NewTicker(interval)
	defer watch.Stop()
	for range watch.C {
		c.mux.Lock()
		for k, v := range c.entry {
			if time.Since(v.createdAt) >= interval {
				delete(c.entry, k)
			}
		}
		c.mux.Unlock()

	}
}
