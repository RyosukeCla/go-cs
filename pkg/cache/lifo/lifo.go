package fifo

import (
	"github.com/RyosukeCla/go-cs/pkg/adt/stack"
	"github.com/RyosukeCla/go-cs/pkg/hashtable/robinhood"
)

type Cache struct {
	cache   robinhood.HashTable
	hits    stack.List
	maxSize int
}

type entry struct {
	key   string
	value interface{}
}

func New(maxSize int) Cache {
	return Cache{
		cache:   robinhood.NewHashTable(uint32(maxSize / 3)),
		hits:    stack.New(),
		maxSize: maxSize,
	}
}

func (c *Cache) Size() int {
	return c.hits.Len()
}

func (c *Cache) Put(key string, value interface{}) {
	c.cache.Put(key, value)
	c.hits.Push(key)
	if c.hits.Len() > c.maxSize {
		lastKey := c.hits.Pop().(string)
		c.cache.Erase(lastKey)
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.cache.Get(key)
}
