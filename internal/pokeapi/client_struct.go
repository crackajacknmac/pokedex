package pokeapi

import (
	"time"
	"github.com/crackajacknmac/pokedex/internal/pokecache"
)

type Client struct {
	cache  pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	cache := pokecache.NewCache(interval)
	return Client {
		cache: cache,
	}
}
