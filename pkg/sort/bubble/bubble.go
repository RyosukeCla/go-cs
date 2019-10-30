package bubble

// Sort sorts array
func Sort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if array[i] < array[j] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
}
