package basic

const PI = 3.14159265358979323846264338327950288419716939937510582097494459
const TWO_PI = 2 * PI
const HALF_PI = PI / 2.0
const E = 2.71828182845904523536028747135266249775724709369995957496696763

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

func Power(a, b float64) float64 {
	// f(x) = x^(1/b) - a = 0 => x = a^b
	//
	return 0
}

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

// Sin returns sin(theta)
func Sin(x float64) float64 {
	xx := PI - Mod(x, TWO_PI)
	res := xx
	preTerm := xx
	for i := 1; i < 8; i++ {
		degree := float64(2*i + 1)
		preTerm = -1 * preTerm * xx * xx / (degree * (degree - 1))
		res += preTerm
	}
	return res
}

// Cos returns sin(theta)
func Cos(x float64) float64 {
	return Sin(HALF_PI - x)
}

// Tan returns tan(theta)
func Tan(x float64) float64 {
	return Sin(x) / Cos(x)
}
