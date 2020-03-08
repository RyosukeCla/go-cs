package interpolation

import (
	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

// Linear returns the values between a and b based on x. And x is in [0,1].
func Linear(a, b, x float64) float64 {
	return a*(1.0-x) + b*x
}

/*
	Cosine returns the values between a and b based on x. And x is in [0,1].

	y(x) := a * (1 - f(x)) + b * f(x), x in [0,1] and f: [0,1] -> [0,1]
	f(x) := (1 - cos(PI * x)) / 2

	y'(0) = y'(1) = 0
	y(0) = a
	y(1) = b
*/
func Cosine(a, b, x float64) float64 {
	x2 := (1.0 - math.Cos(x*math.PI)) / 2.0
	return Linear(a, b, x2)
}
