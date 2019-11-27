package robinhood

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

// HashTable ...
type HashTable struct {
	entries []interface{}
	maxSize int
}

// NewHashTable returns hash table
func NewHashTable(maxSize int) HashTable {
	entries := make([]interface{}, maxSize, maxSize)
	return HashTable{
		entries: entries,
		maxSize: maxSize,
	}
}

func hashFunction(value string) int {
	bytes := []byte(value)
	x := fnv.Hash(bytes)
	if x >= 0 {
		return x
	}
	return -x
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
