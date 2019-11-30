package basic

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	hashtable := NewHashTable(5050)

	for i := 0; i < 100; i++ {
		fmt.Println(i, fmt.Sprintf("%d", i))
		hashtable.Put(fmt.Sprintf("%d", i), i)
	}

	for i := 0; i < 100; i++ {
		if hashtable.Get(fmt.Sprintf("%d", i)) != i {
			t.Fatal("Error")
		}
	}
}
