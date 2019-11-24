package newton

// Args is args for newton method
type Args struct {
	F   func(float64) float64 // F is function
	Fdx func(float64) float64 // Fdx is first order differentiation of F
	N   int                   // N steps
	X0  float64               // initial point of x
}

// Solve finds x s.t. f(x) = 0
// x_{n+1} = x_n - f(x_n)/f'(x_n)
func Solve(args *Args) float64 {
	res := args.X0
	for i := 0; i < args.N; i++ {
		res -= args.F(res) / args.Fdx(res)
	}
	return res
}
