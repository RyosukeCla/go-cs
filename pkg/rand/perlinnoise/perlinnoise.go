package perlinnoise

import (
	"github.com/RyosukeCla/go-cs/pkg/math/interpolation"
	"github.com/RyosukeCla/go-cs/pkg/rand"
	"github.com/RyosukeCla/go-cs/pkg/rand/xorshift"
)

type noise struct {
	noises      []*interpolatedNoise
	persistence float64
	octaves     int
	pre         float64
}

func New(seed uint32, wavelength float64, persistence float64, octaves int) rand.Rand {
	noises := make([]*interpolatedNoise, octaves)
	wavelen := wavelength
	for i := 0; i < octaves; i++ {
		noises[i] = newInterpolatedNoise(wavelen, xorshift.New(uint32(i+1)*seed*1000))
		wavelen /= 2.0
	}

	return &noise{
		noises:      noises,
		persistence: persistence,
		octaves:     octaves,
	}
}

func (n *noise) Generate() float64 {
	total := 0.0
	maxValue := 0.0
	amplitude := 1.0

	for i := 0; i < n.octaves; i++ {
		maxValue += amplitude
		total += n.noises[i].generate() * amplitude
		amplitude *= n.persistence
	}

	return total / maxValue
}

type interpolatedNoise struct {
	rand      rand.Rand
	time      float64
	frequency float64
	y0        float64
	y1        float64
}

func newInterpolatedNoise(wavelength float64, rand rand.Rand) *interpolatedNoise {
	return &interpolatedNoise{
		rand:      rand,
		frequency: 1 / wavelength,
		time:      0,
		y0:        rand.Generate(),
		y1:        rand.Generate(),
	}
}

func (n *interpolatedNoise) generate() float64 {
	n.time += n.frequency

	if n.time >= 1.0 {
		n.time = n.time - 1.0

		// shift
		n.y0 = n.y1
		n.y1 = n.rand.Generate()
	}

	return interpolation.Cosine(n.y0, n.y1, n.time)
}
