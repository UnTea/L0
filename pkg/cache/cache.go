package cache

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewCache() *Cache {
	data := make(map[string]interface{})
	cache := Cache{
		data: data,
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	c.data[key] = value
	c.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()

	item, found := c.data[key]
	if !found {
		return nil, false
	}

	c.RUnlock()

	return item, true
}

func (c *Cache) Delete(key string) {
	c.Lock()
	delete(c.data, key)
	c.Unlock()
}

func (c *Cache) GetAllIDs() []string {
	var ids []string

	c.RLock()

	for key := range c.data {
		ids = append(ids, key)
	}

	c.RUnlock()

	return ids
}
