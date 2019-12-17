package linear

// Rand is for linear congruential generation
type Rand struct {
	previous int
}

// NewRand returns Rand that implements linear congruential generation
func NewRand(seed int) *Rand {
	return &Rand{
		previous: seed,
	}
}

// MODULUS is 2^23 - 1
const MODULUS = 4294967296 - 1

// MULTIPLIER is 48271
const MULTIPLIER = 48271

// INCREMENT is 0
const INCREMENT = 0

// Generate generates rand: X_next = (MULTIPLIER * X_pre + INCREMENT) % MODULUS.
func (r *Rand) Generate() int {
	next := (MULTIPLIER*r.previous + INCREMENT) % MODULUS
	r.previous = next
	return next
}
