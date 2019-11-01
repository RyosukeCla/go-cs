package quick

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestPartition_1(t *testing.T) {
	array := []int{3, 2, 1, 5, 4, 6}
	actual := Partition(array, 0, len(array)-1)
	expected := 2
	if actual != expected {
		t.Fatalf("Error: actual %d, expected %d", actual, expected)
	}
	expectedArray := []int{1, 2, 3, 5, 4, 6}
	if !reflect.DeepEqual(array, expectedArray) {
		t.Fatalf("Error: actual %d, expected %d", array, expectedArray)
	}
}

func TestPartition_2(t *testing.T) {
	array := []int{2, 4, 1, 5}
	actual := Partition(array, 0, len(array)-1)
	expected := 1
	if actual != expected {
		t.Fatalf("Error: actual %d, expected %d", actual, expected)
	}
	expectedArray := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(array, expectedArray) {
		t.Fatalf("Error: actual %d, expected %d", array, expectedArray)
	}
}

func TestSort(t *testing.T) {
	actual := []int{9, 4, 4, 1, 0, 53, 3, 102}

	Sort(actual)

	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error", actual)
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
	Sort(actual)
}
