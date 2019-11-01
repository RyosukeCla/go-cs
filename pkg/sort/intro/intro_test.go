package intro

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	N := 1000
	expected := make([]int, N)
	actual := make([]int, N)

	for i := 0; i < N; i++ {
		random := rand.Intn(10000)
		expected[i] = random
		actual[i] = random
	}

	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Sort(actual)

	for i := 0; i < N-1; i++ {
		if actual[i] > actual[i+1] {
			t.Fatal("Error", i, i+1, actual[i], actual[i+1])
		}
	}
}

func TestSort_2(t *testing.T) {
	N := 1000
	expected := make([]int, N)
	actual := make([]int, N)

	for i := 0; i < N; i++ {
		random := rand.Intn(10000)
		expected[i] = random
		actual[i] = random
	}

	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Sort(actual)

	for i := 0; i < N-1; i++ {
		if actual[i] > actual[i+1] {
			t.Fatal("Error", i, i+1, actual[i], actual[i+1])
		}
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
	Sort(actual)
}
