package pokecache

import (
        "sync"
        "time"
)

type cacheEntry struct {
        createdAt time.Time 
        val []byte
        }

type Cache struct {
        entries map[string]cacheEntry
        mutex sync.Mutex
	interval time.Duration
        }








