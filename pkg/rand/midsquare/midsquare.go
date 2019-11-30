package midsquare

import (
	"fmt"
	"strconv"
)

// Rand for middle square method implementation.
type Rand struct {
	seed int // random seed
	n    int // n-digits
}

// NewRand returns rand
func NewRand(n, seed int) Rand {
	return Rand{
		seed,
		n,
	}
}

// Generate generates pseudorandom. seed is required to be 4 digits number.
func (r *Rand) Generate() int {
	digits := r.seed           // n digits seed
	squared := digits * digits // 2n digits
	str := fmt.Sprintf("%0*d", 2*r.n, squared)
	res, _ := strconv.Atoi(str[r.n/2 : 2*r.n-r.n/2]) // cut middle

	// update seed
	r.seed = res

	return res
}
