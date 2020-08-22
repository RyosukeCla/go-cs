package tlru

import (
	"time"

	"github.com/RyosukeCla/go-cs/pkg/adt/linkedlist"
	"github.com/RyosukeCla/go-cs/pkg/hashtable/robinhood"
)

type Cache struct {
	cache   robinhood.HashTable
	hits    linkedlist.List
	maxSize int
}

type entry struct {
	key       string
	value     interface{}
	timestamp time.Time
}

func New(maxSize int) Cache {
	return Cache{
		cache:   robinhood.NewHashTable(uint32(maxSize / 3)),
		hits:    linkedlist.New(),
		maxSize: maxSize,
	}
}

func (c *Cache) Size() int {
	return c.hits.Len()
}

func (c *Cache) Put(key string, value interface{}, ttl time.Duration) {
	cachedNode := c.cache.Get(key).(*linkedlist.Node)
	if cachedNode != nil {
		c.hits.Remove(cachedNode)
	}

	newNode := c.hits.InsertBiginning(entry{
		key:       key,
		value:     value,
		timestamp: time.Now().Add(ttl),
	})
	c.cache.Put(key, newNode)

	if c.hits.Len() > c.maxSize {
		last := c.hits.Last()
		c.hits.Remove(last)
		c.cache.Erase(last.Value.(entry).key)
	}
}

func (c *Cache) Get(key string) interface{} {
	cachedNode := c.cache.Get(key).(*linkedlist.Node)
	if cachedNode == nil {
		return nil
	}

	if time.Now().After(cachedNode.Value.(entry).timestamp) {
		c.hits.Remove(cachedNode)
		c.cache.Erase(cachedNode.Value.(entry).key)
		return nil
	}

	c.hits.MoveToBeginning(cachedNode)
	return cachedNode.Value.(entry).value
}

func (c *Cache) Prune() {
	head := c.hits.First()
	now := time.Now()
	for {
		if head == nil {
			return
		}
		next := head.Next()

		if now.After(head.Value.(entry).timestamp) {
			c.cache.Erase(head.Value.(entry).key)
			c.hits.Remove(head)
		}

		head = next
	}
}
