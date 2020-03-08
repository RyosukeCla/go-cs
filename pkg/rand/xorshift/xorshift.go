package xorshift

import (
	"github.com/RyosukeCla/go-cs/pkg/rand"
)

const MAX_PLUS_ONE float64 = 4294967296 // 2^32

type Rand struct {
	previous uint32
}

// New returns Rand that implements linear congruential generation
func New(seed uint32) rand.Rand {
	return &Rand{
		previous: seed,
	}
}

// Generate generates rand with xorshift
func (r *Rand) Generate() float64 {
	next := r.previous ^ (r.previous << 13) // (I + Left_A)
	next = next ^ (next >> 17)              // (I + Right_B)
	next = next ^ (next << 5)               // (I + Left_A)
	r.previous = next                       // (I + Left_A)(I + Right_B)(I + Left_A) * bit vector

	return float64(next) / MAX_PLUS_ONE // linear map [0,2^32) to map[0,1)
}
