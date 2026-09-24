package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cacheAdd := cacheEntry{
		createdAt: time.Now(),
		val: val,
		}
	c.entries[key] = cacheAdd
}
