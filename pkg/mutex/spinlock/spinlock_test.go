package spinlock

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	mutex := New()
	runner := func(sig chan int64, sec int64) {
		mutex.Lock()
		defer mutex.Unlock()

		fmt.Println("Locking")
		fmt.Println("sleep", sec, "sec", mutex.lock)
		time.Sleep(time.Duration(sec) * time.Second)
		fmt.Println("Unlock")
		sig <- sec
	}

	sig := make(chan int64, 9)
	defer close(sig)

	for i := 1; i < 10; i++ {
		go runner(sig, int64(i))
	}

	for {
		if len(sig) >= 9 {
			break
		}
	}
}
