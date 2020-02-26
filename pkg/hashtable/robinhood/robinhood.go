package robinhood

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

// HashTable ...
type HashTable struct {
	slots    []slot
	slotSize uint32
}

type entry struct {
	key   string
	value interface{}
}

type slot struct {
	entries []entry
}

// NewHashTable returns hash table
func NewHashTable(slotSize uint32) HashTable {
	slots := make([]slot, slotSize, slotSize)
	size := len(slots)
	for i := 0; i < size; i++ {
		entries := make([]entry, 0, 3)
		slots[i] = slot{
			entries,
		}
	}
	return HashTable{
		slots:    slots,
		slotSize: slotSize,
	}
}

func hashFunction(value string) uint32 {
	bytes := []byte(value)
	return fnv.Hash(bytes)
}

// Get ...
func (h *HashTable) Get(key string) interface{} {
	index := hashFunction(key) % h.slotSize
	slot := &h.slots[index]
	for _, entry := range slot.entries {
		if entry.key == key {
			return entry.value
		}
	}
	return nil
}

// Put ...
func (h *HashTable) Put(key string, value interface{}) {
	index := hashFunction(key) % h.slotSize
	slot := &h.slots[index]

	for _, entry := range slot.entries {
		if entry.key == key {
			entry.value = value
			return
		}
	}

	slot.entries = append(slot.entries, entry{
		key,
		value,
	})
}

// Erase ...
func (h *HashTable) Erase(key string) {
	index := hashFunction(key) % h.slotSize
	slot := &h.slots[index]
	for i, entry := range slot.entries {
		if entry.key == key {
			entries := slot.entries
			entries[0], entries[i] = entries[i], entries[0] // swap
			slot.entries = entries[1:]                      // remove
			return
		}
	}
}

func (h *HashTable) Size() int {
	size := 0
	for _, slot := range h.slots {
		size += len(slot.entries)
	}
	return size
}
