package bucket

// Sort sorts array. for every element of an array must be in [0, upper).
func Sort(array []int, upper int) []int {
	N := len(array)
	counts := make([]int, upper)
	for _, element := range array {
		counts[element]++
	}

	res := make([]int, N, N)
	index := 0
	for i, count := range counts {
		for j := 0; j < count; j++ {
			res[index] = i
			index++
		}
	}

	return res
}
