package pokecache

import (
"time"
//"sync"
)

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]cacheEntry)
	c := &Cache {
		entries: entries,
		//mutex: sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}
