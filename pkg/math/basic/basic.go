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
	} else {
		return -x
	}
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
