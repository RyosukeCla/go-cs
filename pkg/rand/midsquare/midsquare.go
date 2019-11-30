package midsquare

import (
	"fmt"
	"strconv"
)

// Rand ...
type Rand struct {
	seed int
	n    int
}

// NewRand returns psedou random with n-digits.
func NewRand(n int, seed int) Rand {
	return Rand{
		seed: seed,
		n:    n,
	}
}

// Generate generates pseudorandom. seed is required to be 4 digits number.
func (rand *Rand) Generate() int {
	digits := rand.seed        // n digits seed
	squared := digits * digits // 2n digits
	str := fmt.Sprintf("%0*d", 2*rand.n, squared)
	res, _ := strconv.Atoi(str[rand.n/2 : 2*rand.n-rand.n/2]) // cut middle

	// update seed
	rand.seed = res

	return res
}
