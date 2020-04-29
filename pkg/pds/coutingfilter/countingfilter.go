package countingfilter

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

const MAX = 4294967295

type CountingFilter struct {
	counters []uint32
	m        int // bits size
	k        int // hash count
}

func New(m, k int) *CountingFilter {
	counters := make([]uint32, m)
	return &CountingFilter{
		counters: counters,
		m:        m,
		k:        k,
	}
}

func (bf *CountingFilter) Add(key []byte) {
	for i := 0; i < bf.k; i++ {
		index := hash(key, i) % bf.m
		if bf.counters[index] <= MAX {
			bf.counters[index]++
		}
	}
}

func (bf *CountingFilter) Delete(key []byte) {
	if !bf.Check(key) {
		return
	}
	for i := 0; i < bf.k; i++ {
		index := hash(key, i) % bf.m
		bf.counters[index]--
	}
}

func (bf *CountingFilter) Check(key []byte) bool {
	for i := 0; i < bf.k; i++ {
		index := hash(key, i) % bf.m
		if bf.counters[index] == 0 {
			return false
		}
	}
	return true
}

func hash(input []byte, i int) int {
	return int(fnv.Hash([]byte(string(i*23456789) + string(input))))
}
