package complex

import (
	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

type Complex struct {
	Re float64
	Im float64
}

func New(re, im float64) *Complex {
	return &Complex{
		Re: re,
		Im: im,
	}
}

func NewRe(re float64) *Complex {
	return New(re, 0)
}

func NewIm(im float64) *Complex {
	return New(0, im)
}

// Euler returns e^z
func Euler(z *Complex) *Complex {
	cf := math.Exp(z.Re)

	return &Complex{
		Re: cf * math.Cos(z.Im),
		Im: cf * math.Sin(z.Im),
	}
}

// Abs returns |z|
func Abs(z *Complex) float64 {
	return math.Sqrt(z.Re * z.Re + z.Im * z.Im)
}

// Arg returns arg(z): phi part of polar form
func Arg(z *Complex) float64 {
	if z.Im != 0 {
		return 2 * math.Arctan((math.Sqrt(z.Re * z.Re + z.Im * z.Im) - z.Re) / z.Im)
	} else if z.Re > 0 && z.Im == 0 {
		return 0
	} else if z.Re < 0 && z.Im == 0 {
		return math.PI
	} else {
		return 0 // undefined
	}
}

// Add returns a + b
func Add(a *Complex, b *Complex) *Complex {
	return &Complex{
		Re: a.Re + b.Re,
		Im: a.Im + b.Im,
	}
}

// Sub returns a - b
func Sub(a *Complex, b *Complex) *Complex {
	return &Complex{
		Re: a.Re - b.Re,
		Im: a.Im - b.Im,
	}
}

// Mul returns a * b
func Mul(a *Complex, b *Complex) *Complex {
	return &Complex{
		Re: a.Re * b.Re - a.Im * b.Im,
		Im: a.Re * b.Im + a.Im * b.Re,
	}
}

// Reciprocal returns 1 / a
func Reciprocal(a *Complex) *Complex {
	cf := 1 / (a.Re * a.Re + a.Im + a.Im)
	return &Complex{
		Re: cf * a.Re,
		Im: - cf * a.Im,
	}
}

// Div returns a / b
func Div(a *Complex, b *Complex) *Complex {
	return Mul(a, Reciprocal(b))
}

// Conjugate returns conjugate of a
func Conjugate(a *Complex) *Complex {
	return &Complex{
		Re: a.Re,
		Im: -a.Im,
	}
}

// Equal returns equality of two complex numbers
func Equal(a *Complex, b *Complex) bool {
	if a.Re != b.Re || a.Im != a.Im {
		return false
	}
	return true
}

// AlmostEqual is epsilon comparison
func AlmostEqual(a *Complex, b *Complex) bool {
	if !math.AlmostEqual(a.Re, b.Re) || !math.AlmostEqual(a.Im, b.Im) {
		return false
	}
	return true
}