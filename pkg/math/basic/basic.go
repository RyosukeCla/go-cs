package basic

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

// Ceil returns upper [x]
func Ceil(x float64) float64 {
	floor := float64(int64(x))
	if floor < x {
		return floor + 1
	}
	return floor
}

// Sin returns sin(theta)
func Sin(x float64) float64 {
	res := x
	preTerm := x
	for i := 1; i < 7; i++ {
		degree := float64(2*i + 1)
		preTerm = -1 * preTerm * x * x / (degree * (degree - 1))
		res += preTerm
	}
	return res
}

// Cos returns sin(theta)
func Cos(x float64) float64 {
	res := x
	preTerm := 1.0
	for i := 1; i < 7; i++ {
		degree := float64(2 * i)
		preTerm = -1 * preTerm * x * x / (degree * (degree - 1)) // nextTerm = -1 * preTerm * x^2 / (degree * (degree - 1))
		res += preTerm
	}
	return res
}

// Power returns x^y, and y is integer
// func Power(x float64, y int) float64 {
// 	for {

// 	}
// 	return
// }

// // Sin returns sin(x)
// func Sin(x float64) float64 {
// 	res := x
// 	N := 10

// 	for n := 0; n < N; n++ {
// 		res +=
// 	}
// }
