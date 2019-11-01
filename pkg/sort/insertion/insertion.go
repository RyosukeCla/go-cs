package insertion

// Sort sorts an array with insertion sorting algorithm
func Sort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j--
		}
	}
}
