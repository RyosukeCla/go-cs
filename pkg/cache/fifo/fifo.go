package fifo

import (
	"github.com/RyosukeCla/go-cs/pkg/adt/queue"
	"github.com/RyosukeCla/go-cs/pkg/hashtable/robinhood"
)

type Cache struct {
	cache   robinhood.HashTable
	hits    queue.List
	maxSize int
}

type entry struct {
	key   string
	value interface{}
}

func New(maxSize int) Cache {
	return Cache{
		cache:   robinhood.NewHashTable(uint32(maxSize / 3)),
		hits:    queue.New(),
		maxSize: maxSize,
	}
}

func (c *Cache) Size() int {
	return c.hits.Len()
}

func (c *Cache) Put(key string, value interface{}) {
	c.cache.Put(key, value)
	c.hits.Enqueue(key)
	if c.hits.Len() > c.maxSize {
		firstKey := c.hits.Dequeue().(string)
		c.cache.Erase(firstKey)
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.cache.Get(key)
}
