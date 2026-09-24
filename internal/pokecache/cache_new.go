package pokecache

import (
"time"
"sync"
)

func NewCache(interval time.Duration) Cache {
	entries := make(map[string]cacheEntry)
	return Cache {
		entries: entries,
		mutex: sync.Mutex{},
		interval: interval,
	}

}
