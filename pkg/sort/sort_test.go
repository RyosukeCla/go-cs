package sort

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

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
		{name: "Bubble", sort: bubble.Sort},
		{name: "Heap", sort: heap.Sort},
		{name: "Insertion", sort: insertion.Sort},
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

func TestCompareQuickWithMerge(t *testing.T) {
	n := 20000
	divs := 1000
	size := 20000
	mergeSortTimes := make([]int, n, n)
	quickSortTimes := make([]int, n, n)

	// monte calro
	for i := 0; i < n; i++ {
		array := makeRandomArray(size)
		current := time.Now()
		merge.Sort(array)
		mergeSortTimes[i] = int(time.Since(current))
	}

	for i := 0; i < n; i++ {
		array := makeRandomArray(size)
		current := time.Now()
		quick.Sort(array)
		quickSortTimes[i] = int(time.Since(current))
	}

	// make histgrams
	quick.Sort(mergeSortTimes)
	quick.Sort(quickSortTimes)
	min := int(math.Min(float64(quickSortTimes[0]), float64(mergeSortTimes[0])))
	max := int(math.Max(float64(quickSortTimes[n-1]), float64(mergeSortTimes[n-1])))
	rang := max - min
	delta := int(rang / divs)

	mergeDist := make([]int, divs)
	quickDist := make([]int, divs)

	for i := 0; i < divs; i++ {
		mergeDist[i] = 0
		quickDist[i] = 0
	}

	currentTime := min
	index := 0
	index2 := 0
	for {
		if index >= n || index2 >= divs {
			break
		}
		if mergeSortTimes[index] < currentTime+delta {
			mergeDist[index2]++
			index++
		} else {
			currentTime += delta
			index2++
		}
	}

	currentTime = min
	index = 0
	index2 = 0
	for {
		if index >= n || index2 >= divs {
			break
		}
		if quickSortTimes[index] < currentTime+delta {
			quickDist[index2]++
			index++
		} else {
			currentTime += delta
			index2++
		}
	}

	// calc prob
	count := 0
	for i := divs - 1; i >= 0; i-- {
		if quickDist[i] > mergeDist[i] {
			count += quickDist[i]
		}
	}

	fmt.Println("Prob(quick > merge) =", float64(count)/float64(n))
}
