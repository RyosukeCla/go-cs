package midsquare

import (
	"testing"

	"github.com/RyosukeCla/go-cs/pkg/hashtable/robinhood"
)

func Test(t *testing.T) {
	N := 10
	seen := robinhood.NewHashTable(uint(5000))
	rand := NewRand(4, 9090)
	for i := 0; i < N; i++ {
		r := rand.Generate()
		if seen.Get(string(r)) != nil {
			t.Fatal("error")
		}
		seen.Put(string(r), r)
	}
}
