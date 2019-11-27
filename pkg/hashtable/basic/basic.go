package basic

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

// HashTable ...
type HashTable struct {
	entries []interface{}
	maxSize uint
}

// NewHashTable returns hash table
func NewHashTable(maxSize uint) HashTable {
	entries := make([]interface{}, maxSize, maxSize)
	return HashTable{
		entries: entries,
		maxSize: maxSize,
	}
}

func hashFunction(value string) uint {
	bytes := []byte(value)
	return fnv.Hash(bytes)
}

// Get ...
func (h *HashTable) Get(key string) interface{} {
	index := hashFunction(key) % h.maxSize
	return h.entries[index]
}

// Put ...
func (h *HashTable) Put(key string, entry interface{}) {
	index := hashFunction(key) % h.maxSize
	h.entries[index] = entry
}

// Erase ...
func (h *HashTable) Erase(key string) {
	index := hashFunction(key) % h.maxSize
	h.entries[index] = nil
}
