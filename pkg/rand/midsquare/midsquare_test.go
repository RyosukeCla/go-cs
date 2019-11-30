package midsquare

import (
	"testing"

	"github.com/RyosukeCla/go-cs/pkg/hashtable/basic"
)

func Test(t *testing.T) {
	N := 10
	seen := basic.NewHashTable(uint(5050))
	rand := NewRand(4, 9090)
	for i := 0; i < N; i++ {
		r := rand.Generate()
		if seen.Get(string(r)) != nil {
			t.Fatal("error")
		}
		seen.Put(string(r), r)
	}
}
