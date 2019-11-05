package sort

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/RyosukeCla/go-cs/pkg/sort/bubble"
	"github.com/RyosukeCla/go-cs/pkg/sort/heap"
	"github.com/RyosukeCla/go-cs/pkg/sort/insertion"
	"github.com/RyosukeCla/go-cs/pkg/sort/intro"
	"github.com/RyosukeCla/go-cs/pkg/sort/merge"
	"github.com/RyosukeCla/go-cs/pkg/sort/quick"
)

func makeRandomArray(N int) []int {
	array := make([]int, N)
	for i := 0; i < N; i++ {
		random := rand.Intn(10000)
		array[i] = random
	}
	return array
}

func bench(b *testing.B, sortFunc func(array []int), size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		array := makeRandomArray(size)
		b.StartTimer()
		sortFunc(array)
		b.StopTimer()
	}
}

type sortAlgorithms struct {
	name string
	sort func(array []int)
}

func BenchmarkSorts(b *testing.B) {
	sortAlgorithms := []sortAlgorithms{
		{name: "Insertion", sort: insertion.Sort},
		{name: "Heap", sort: heap.Sort},
		{name: "Bubble", sort: bubble.Sort},
		{name: "Quick", sort: quick.Sort},
		{name: "Merge", sort: merge.Sort},
		{name: "Intro", sort: intro.Sort},
	}

	sizes := []int{1000, 10 * 1000, 100 * 1000}

	for _, sortAlgorithm := range sortAlgorithms {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%s_%d", sortAlgorithm.name, size), func(b *testing.B) {
				bench(b, sortAlgorithm.sort, size)
			})
		}
	}
}
