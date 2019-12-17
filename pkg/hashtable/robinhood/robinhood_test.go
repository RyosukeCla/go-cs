package robinhood

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	hashtable := NewHashTable(100)

	for i := 0; i < 200; i++ {
		hashtable.Put(fmt.Sprintf("%d", i), i)
	}

	for i := 0; i < 200; i++ {
		if hashtable.Get(fmt.Sprintf("%d", i)) != i {
			t.Fatal("Error")
		}
	}

	hashtable.Erase("33")
	if hashtable.Get("33") != nil {
		t.Fatal("Error")
	}
}
