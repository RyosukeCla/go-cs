package greedy

// Find s s.t. g^s = y and y, g in Z mod p
func Find(y, g, p int) int {
	s := 0
	gs := 1
	for {
		s++
		gs *= g
		gs %= p

		if gs == y {
			return s
		}
	}
}
