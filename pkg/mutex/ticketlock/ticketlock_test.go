package ticketlock

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	mutex := New()
	runner := func(sig chan int64, sec int64) {
		mutex.Acquire()
		defer mutex.Release()

		fmt.Println("Locking")
		fmt.Println("sleep", sec, "mili sec")
		time.Sleep(time.Duration(sec) * time.Millisecond)
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
