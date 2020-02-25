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

func New(cap int) cache {
	lock := ticketlock.New()
	return cache{
		store:   newStore(cap),
		backing: newBackingStore(cap),
		mutex:   &lock,
	}
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
	c.store.Write(data)
	go c.backing.Write(data)
}

func (c *cache) WriteAround(data int) {
	c.backing.Write(data)
}

func (c *cache) Read(index int) int {
	return c.store.Read(index)
}

func (c *cache) ReadFromBackingStore(index int) int {
	return c.backing.Read(index)
}

type store struct {
	data []int
}

func newStore(cap int) *store {
	return &store{
		data: make([]int, 0, cap),
	}
}

func (c *store) Write(data int) {
	c.data = append(c.data, data)
}

func (c *store) Read(index int) int {
	return c.data[index]
}

type backgingstore struct {
	data  []int
	mutex *ticketlock.Mutex
}

func newBackingStore(cap int) *backgingstore {
	lock := ticketlock.New()
	return &backgingstore{
		data:  make([]int, 0, cap),
		mutex: &lock,
	}
}

func (c *backgingstore) Write(data int) {
	time.Sleep(100) // disk i/o is slow
	c.data = append(c.data, data)
}

func (c *backgingstore) Read(index int) int {
	time.Sleep(100) // disk io is slow
	return c.data[index]
}
