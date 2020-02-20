package write

import (
	"time"

	"github.com/RyosukeCla/go-cs/pkg/mutex/ticketlock"
)

type cache struct {
	store   *store
	backing *backgingstore
	mutex   *ticketlock.Mutex
}

func (c *cache) Write(data int) {
	c.store.Write(data)
}

func (c *cache) WriteThrough(data int) {
	c.mutex.Acquire()
	defer c.mutex.Release()

	c.store.Write(data)
	c.backing.Write(data)
}

func (c *cache) WriteBack(data int) {
	c.mutex.Acquire()
	defer c.mutex.Release()

	c.store.Write(data)
	go c.backing.Write(data)
}

func (c *cache) WriteAround(data int) {
	c.backing.Write(data)
}

func (c *cache) Read(index int) int {
	return c.store.Read(index)
}

type store struct {
	data []int
}

func (c *store) Write(data int) {
	c.data = append(c.data, data)
}

func (c *store) Read(index int) int {
	return c.data[index]
}

type backgingstore struct {
	data []int
}

func (c *backgingstore) Write(data int) {
	time.Sleep(2) // disk i/o is slow
	c.data = append(c.data, data)
}

func (c *backgingstore) Read(index int) int {
	time.Sleep(2)
	return c.data[index]
}
