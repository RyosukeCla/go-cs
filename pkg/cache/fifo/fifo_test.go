package fifo

import (
	"testing"
)

func Test(t *testing.T) {
	cache := New(3)
	cache.Put("1", 0)
	cache.Put("2", 0)
	cache.Put("3", 0)
	if cache.Get("1") != 0 {
		t.Fatal("error")
	}
	cache.Put("4", 0)
	if cache.Get("1") != nil {
		t.Fatal("error")
	}
}
