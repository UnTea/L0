package cache

import (
	"sync"
)

// Cache - represents cache
type Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewCache - creates new instance of cache
func NewCache() *Cache {
	data := make(map[string]interface{})

	cache := Cache{
		data: data,
	}

	return &cache
}

// Set - adds new data in cache
func (c *Cache) Set(key string, value interface{}) {
	c.Lock()

	c.data[key] = value

	c.Unlock()
}

// Get - gets data from cache by id
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()

	item, found := c.data[key]
	if !found {
		return nil, false
	}

	c.RUnlock()

	return item, true
}

// Delete - deletes data from cache
func (c *Cache) Delete(key string) {
	c.Lock()

	delete(c.data, key)

	c.Unlock()
}

// GetAllIDs - gets all IDs from cache
func (c *Cache) GetAllIDs() []string {
	var ids []string

	c.RLock()

	for key := range c.data {
		ids = append(ids, key)
	}

	c.RUnlock()

	return ids
}
