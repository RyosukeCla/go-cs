package bloomfilter

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

type BloomFilter struct {
	bits []bool
	m    int // bits size
	k    int // hash count
}

func New(m, k int) *BloomFilter {
	bits := make([]bool, m)
	return &BloomFilter{
		bits: bits,
		m:    m,
		k:    k,
	}
}

func (bf *BloomFilter) Add(key []byte) {
	for i := 0; i < bf.k; i++ {
		index := hash(key, i) % bf.m
		bf.bits[index] = true
	}
}

func (bf *BloomFilter) Check(key []byte) bool {
	for i := 0; i < bf.k; i++ {
		index := hash(key, i) % bf.m
		if bf.bits[index] == false {
			return false
		}
	}
	return true
}

func hash(input []byte, i int) int {
	return int(fnv.Hash([]byte(string(i*23456789) + string(input))))
}
