package midsquare

import (
	"fmt"
	"strconv"

	"github.com/RyosukeCla/go-cs/pkg/rand"
)

const MAX_PLUS_ONE float64 = 10000

// Rand for middle square method implementation.
type Rand struct {
	seed int // random seed
	n    int // n-digits
}

// NewRand returns rand
func New(n, seed int) rand.Rand {
	return &Rand{
		seed,
		n,
	}
}

// Generate generates pseudorandom. seed is required to be 4 digits number.
func (r *Rand) Generate() float64 {
	digits := r.seed           // n digits seed
	squared := digits * digits // 2n digits
	str := fmt.Sprintf("%0*d", 2*r.n, squared)
	res, _ := strconv.Atoi(str[r.n/2 : 2*r.n-r.n/2]) // cut middle

	// update seed
	r.seed = res

	return float64(res) / MAX_PLUS_ONE // map [0,10000) to [0,1)
}
