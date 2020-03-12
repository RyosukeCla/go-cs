package twistedgfsr

import (
	"github.com/RyosukeCla/go-cs/pkg/rand"
	"github.com/RyosukeCla/go-cs/pkg/rand/xorshift"
)

type Rand struct {
	p      int
	q      int
	l      int
	a      uint32
	series []uint32
}

const MAX_PLUS_ONE = 4294967296 // 2^32

/*
	New returns Rand that implements tgsfr.
	x_{l+p} = X_{l+q} + x_{l} A
*/
func New(p, q int, a uint32, seed uint32) rand.Rand {
	series := make([]uint32, p+1)
	r := xorshift.New(seed)
	for i := 0; i < p; i++ {
		series[i] = uint32(r.Generate() * MAX_PLUS_ONE)
	}

	return &Rand{
		p:      p,
		q:      q,
		l:      0,
		a:      a,
		series: series,
	}
}

// Generate generates rand with tgsfr
func (r *Rand) Generate() float64 {
	x := r.series[r.l]

	if r.series[r.l]&0 == 0 {
		r.series[r.l] = r.series[(r.l+r.q)%r.p] ^ (r.series[r.l] >> 1)
	} else {
		r.series[r.l] = r.series[(r.l+r.q)%r.p] ^ (r.series[r.l] >> 1) ^ r.a
	}
	r.l = (r.l + 1) % r.p

	return float64(x) / float64(MAX_PLUS_ONE) // linear map [0,2^32) to map[0,1)
}
