package adt

// MaxHeap is max heap
type MaxHeap struct {
	array []float32
}

// NewMaxHeap returns max heap
func NewMaxHeap(cap int) MaxHeap {
	array := make([]float32, 0, cap)
	return MaxHeap{
		array: array,
	}
}

// Insert inserts element with keeping max heap property
func (h *MaxHeap) Insert(element float32) {
	h.array = append(h.array, element)
	current := len(h.array) - 1
	for {
		parent := parentIndex(current)
		if h.array[current] > h.array[parent] {
			h.array[current], h.array[parent] = h.array[parent], h.array[current]
			current = parent
		} else {
			break
		}
	}
}

// Extract extracts max element with keeping max heap property
func (h *MaxHeap) Extract() float32 {
	max := h.array[0]
	maxHeapify(h.array, 0, len(h.array))
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-2]
	return max
}

// MinHeap is min heap
type MinHeap struct {
	array []float32
}

// NewMinHeap returns min heap
func NewMinHeap(cap int) MinHeap {
	array := make([]float32, 0, cap)
	return MinHeap{
		array: array,
	}
}

// Insert inserts element with keeping min heap property
func (h *MinHeap) Insert(element float32) {
	h.array = append(h.array, element)
	current := len(h.array) - 1
	for {
		parent := parentIndex(current)
		if h.array[current] < h.array[parent] {
			h.array[current], h.array[parent] = h.array[parent], h.array[current]
			current = parent
		} else {
			break
		}
	}
}

// Extract extracts min element with keeping min heap property
func (h *MinHeap) Extract() float32 {
	min := h.array[0]
	minHeapify(h.array, 0, len(h.array))
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-2]
	return min
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func maxHeapify(arr []float32, i int, size int) {
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	largest := i

	// If left child is larger than root
	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		// swap
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		maxHeapify(arr, largest, size)
	}
}

func minHeapify(arr []float32, i int, size int) {
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	largest := i

	// If left child is larger than root
	if left < size && arr[left] < arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < size && arr[right] < arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		// swap
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		minHeapify(arr, largest, size)
	}
}
