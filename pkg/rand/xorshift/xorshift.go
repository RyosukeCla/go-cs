package xorshift

type Rand struct {
	previous uint32
}

// NewRand returns Rand that implements linear congruential generation
func NewRand(seed uint32) *Rand {
	return &Rand{
		previous: seed,
	}
}

// Generate generates rand with xorshift
func (r *Rand) Generate() uint32 {
	next := r.previous ^ (r.previous << 13) // (I + Left_A)
	next = next ^ (next >> 17)              // (I + Right_B)
	next = next ^ (next << 5)               // (I + Left_A)
	r.previous = next
	return next // (I + Left_A)(I + Right_B)(I + Left_A) * bit vector
}
