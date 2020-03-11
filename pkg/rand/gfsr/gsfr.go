package gsfr

import (
	"github.com/RyosukeCla/go-cs/pkg/rand"
	"github.com/RyosukeCla/go-cs/pkg/rand/xorshift"
)

type Rand struct {
	p      int
	q      int
	k      int
	j      int
	series []uint32
}

const MAX_PLUS_ONE = 4294967296 // 2^32

/*
	New returns Rand that implements gsfr.
	p and q (p > q) have to be selected to satisfy primitive polynomials, such that
		x^p + x^q + 1.
	And period is 2^p - 1.

	e.g.) (p, q) in {(47, 21), (95, 17), (111, 45), ...}

	further detail:
		Lewis, T. G., & Payne, W. H. Generalized feedback shift register pseudorandom number generator. Journal of the Association for Computing Machinery, 1973, in press.
*/
func New(p, q int, seed uint32) rand.Rand {
	series := make([]uint32, p+1)
	r := xorshift.New(seed)
	for i := 0; i < p; i++ {
		series[i] = uint32(r.Generate() * MAX_PLUS_ONE)
	}

	return &Rand{
		p:      p,
		q:      q,
		k:      0,
		j:      0,
		series: series,
	}
}

// Generate generates rand with gsfr
func (r *Rand) Generate() float64 {
	r.k++          // k <- k + 1
	if r.k > r.p { // if k > p, set k <- 1
		r.k = 1
	}
	r.j = r.k + r.q // j <- k + q
	if r.j > r.p {  // if j > p, set j <- j - p
		r.j -= r.p
	}

	x := r.series[r.k] ^ r.series[r.j] // S_k xor S_j
	r.series[r.k] = x                  // store S_k <- S_k xor S_j

	return float64(x) / float64(MAX_PLUS_ONE) // linear map [0,2^32) to map[0,1)
}
