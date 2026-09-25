package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}
