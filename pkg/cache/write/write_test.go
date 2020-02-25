package write

import (
	"testing"
)

func TestWriteThrough(t *testing.T) {
	c := New(10)

	runner := func(sig chan int, i int) {
		c.WriteThrough(i)
		sig <- i
	}

	sig := make(chan int, 10)
	defer close(sig)
	for i := 0; i < 10; i++ {
		go runner(sig, i)
	}

	for {
		if len(sig) >= 10 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if c.Read(i) != c.ReadFromBackingStore(i) {
			t.Fatal("error")
		}
	}
}

func BenchmarkWriteThrough(b *testing.B) {
	size := 2000
	c := New(size)

	runner := func(sig chan int, i int) {
		c.WriteThrough(i)
		sig <- i
	}

	sig := make(chan int, size)
	defer close(sig)

	b.ResetTimer()
	for i := 0; i < size; i++ {
		go runner(sig, i)
	}

	for {
		if len(sig) >= size {
			break
		}
	}
	b.StopTimer()
}

func BenchmarkWriteBack(b *testing.B) {
	size := 2000
	c := New(size)

	runner := func(sig chan int, i int) {
		c.WriteBack(i)
		sig <- i
	}

	sig := make(chan int, size)
	defer close(sig)

	b.ResetTimer()
	for i := 0; i < size; i++ {
		go runner(sig, i)
	}

	for {
		if len(sig) >= size {
			break
		}
	}
	b.StopTimer()
}
