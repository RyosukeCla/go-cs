package basic

import (
	"testing"
)

func TestHashtable(t *testing.T) {
	hashtable := NewHashTable(100)
	hashtable.Put("key", "value")
	value := hashtable.Get("key")

	if value != "value" {
		t.Fatal("Error")
	}
}
