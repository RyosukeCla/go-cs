package consistent

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	hashtable := NewHashTable([]string{})

	for i := 0; i < 200; i++ {
		hashtable.AddNode(fmt.Sprintf("server-%d", i))
	}

	for i := 0; i < 500; i++ {
		if hashtable.Get(fmt.Sprintf("key-%d", i)) == nil {
			t.Fatal("Error")
		}
	}
}
