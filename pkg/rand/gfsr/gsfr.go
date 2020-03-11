/*
reference

	W. H. PAYNE. "Pseudorandom numbers for mini- and microcomputers: A generalized feedback shift register algorithm" Behav. Res.l\Ieth. & Instru.,1973, Vol. 5(2)
*/

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

// New returns Rand that implements linear congruential generation
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
	r.k++
	if r.k > r.p {
		r.k = 1
	}
	r.j = r.k + r.q
	if r.j > r.p {
		r.j -= r.p
	}

	x := r.series[r.k] ^ r.series[r.j]
	r.series[r.k] = x

	return float64(x) / float64(MAX_PLUS_ONE) // linear map [0,2^32) to map[0,1)
}
