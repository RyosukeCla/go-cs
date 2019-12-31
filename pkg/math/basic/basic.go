package basic

const PI float64 = 3.14159265358979323846264338327950288419716939937510582097494459
const TWO_PI float64 = 2 * PI
const HALF_PI float64 = PI / 2.0
const E float64 = 2.71828182845904523536028747135266249775724709369995957496696763

// Factorial returns n!
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Abs returns absolute value of x
func Abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -x
}

// Square returs square of x
func Square(x float64) float64 {
	return x * x
}

// Mod returns r s.t. a / b = c ... r, c in Z.
func Mod(a, b float64) float64 {
	ab := a / b
	return a - b*Floor(ab)
}

// func Power(a, b float64) float64 {
// 	// f(x) = x^(1/b) - a = 0 => x = a^b
// 	//

// 	return Exp(y * Log(x))
// 	return 0
// }

// Ceil returns upper [x]
func Ceil(x float64) float64 {
	floor := float64(int64(x))
	if floor < x {
		return floor + 1
	}
	return floor
}

// Floor returns lower [x]
func Floor(x float64) float64 {
	y := float64(int64(x))
	if y < 0 {
		return y - 1
	}
	return y
}

// Exp returns e^x = 1 + x + x^2 / 2! + ...
func Exp(x float64) float64 {
	res := 1 + x
	preTerm := x
	for i := 2; i < 18; i++ {
		degree := float64(i)
		preTerm = preTerm * x / degree
		res += preTerm
	}
	return res
}

// Sin returns sin(x)
func Sin(x float64) float64 {
	xx := PI - Mod(x, TWO_PI) // translate to -PI <= x < PI
	res := xx
	preTerm := xx
	for i := 1; i < 10; i++ {
		degree := float64(2*i + 1)
		preTerm = -1 * preTerm * xx * xx / (degree * (degree - 1))
		res += preTerm
	}
	return res
}

// Cos returns cos(x) = sin(PI/2 - x)
func Cos(x float64) float64 {
	return Sin(HALF_PI - x)
}

// Tan returns tan(x) = sin(x) / cos(x)
func Tan(x float64) float64 {
	return Sin(x) / Cos(x)
}
