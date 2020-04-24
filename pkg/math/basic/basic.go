package basic

const PI float64 = 3.14159265358979323846264338327950288419716939937510582097494459
const TWO_PI float64 = 2 * PI
const HALF_PI float64 = PI / 2.0
const E float64 = 2.71828182845904523536028747135266249775724709369995957496696763
const LN2 float64 = 0.693147180559945309417232121458176568075500134360255254120680009
const LN10 float64 = 2.30258509299404568401799145468436420760110148862877297603332790

// Factorial returns n!
func Factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Perm returns a P b (Permutation)
func Perm(a, b int64) int64 {
	res := int64(1)
	for i := a; i >= a-b+1; i-- {
		res *= i
	}

	return res
}

// Comb returns a C b (Combination)
func Comb(a, b int64) int64 {
	return Perm(a, b) / Factorial(b)
}

// BinoimialCf is alias of Comb (Binoimial Coefficient)
func BinoimialCf(a, b int64) int64 {
	return Comb(a, b)
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

// Mod returns r s.t. a / b = c + r, c in Z.
func Mod(a, b float64) float64 {
	ab := a / b
	return a - b*Floor(ab)
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
	return -Ceil(-x)
}

// Sqrt returns square root of x
func Sqrt(x float64) float64 {
	res := 1.0
	for i := 0; i < 7; i++ {
		res -= (res*res - x) / (2 * res) // newton method
	}
	return res
}

// Power returns a^b
func Power(a, b float64) float64 {
	return Exp(b * Ln(a))
}

// ln(x) = 2 arctanh (x-1)/(x+1)
func lnWithTaylorExpansion(x float64) float64 {
	zz := (x - 1) / (x + 1)
	preTerm := zz
	res := preTerm
	for i := 1; i < 8; i++ {
		degree := float64(2*i + 1)
		preTerm = preTerm * zz * zz
		res += preTerm / degree
	}
	return 2 * res
}

// Ln returns log_e(x)
func Ln(x float64) float64 {
	// find a and b s.t. x = a * 2^b, 0 < a < 1.0
	// then calculate ln(x) = ln(a) + b * ln2
	a := x
	b := 0.0
	for {
		if a < 1.0 {
			break
		}
		a = a * 0.5
		b++
	}

	return lnWithTaylorExpansion(a) + b*LN2
}

// Log returns log_a(b)
func Log(a, b float64) float64 {
	return Ln(b) / Ln(a)
}

func expWithTalyorExpansion(x float64) float64 {
	res := 1 + x
	preTerm := x
	for i := 2; i < 8; i++ {
		degree := float64(i)
		preTerm = preTerm * x / degree
		res += preTerm
	}
	return res
}

// Exp returns e^x
func Exp(x float64) float64 {
	// find c and r s.t. x = c + r, c in Z and r in [0, 1)
	// then calculate e^x = e^c * e^r
	c := Floor(x)
	r := x - c
	res := 1.0
	// e^c
	if c > 0 {
		n := int(c)
		for i := 0; i < n; i++ {
			res *= E
		}
	} else {
		n := int(-c)
		inverse := 1.0 / E
		for i := 0; i < n; i++ {
			res *= inverse
		}
	}
	// * e^r
	return res * expWithTalyorExpansion(r)
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

// Arcsin returns arcsin(x) x in [-1,1]
func Arcsin(x float64) float64 {
	res := x
	preX := x
	for i := 1; i < 10; i++ {
		preX = preX * x * x
		cf := 1 / Power(4.0, float64(i)) * float64(BinoimialCf(int64(2*i), int64(i))) / float64(2*i+1)
		res += cf * preX
	}
	return res
}

// Arccos returns arccos(x) x in [-1,1]
func Arccos(x float64) float64 {
	return HALF_PI - Arcsin(x)
}

// Arctan returns arctan(x) x in [-1,1]
func Arctan(x float64) float64 {
	return Arcsin(x / Sqrt(x*x+1))
}
