package bucket

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	array := []int{9, 4, 4, 1, 0, 53, 3, 102}

	actual := Sort(array, 103)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error")
	}
}

func BenchmarkSort(b *testing.B) {
	N := b.N
	expected := make([]int, N)
	actual := make([]int, N)

	for i := 0; i < N; i++ {
		random := rand.Intn(10000)
		expected[i] = random
		actual[i] = random
	}

	b.ResetTimer()
	Sort(actual, 10000)
}
