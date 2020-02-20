package counting

import (
	"runtime"
	"sync/atomic"
)

type semaphore struct {
	value int32
}

func New() semaphore {
	return semaphore{
		value: int32(1),
	}
}

func (s *semaphore) Acquire() {
	for {
		if s.value > 0 {
			atomic.AddInt32(&s.value, -1)
			return
		}
		runtime.Gosched()
	}
}

func (s *semaphore) Release() {
	atomic.AddInt32(&s.value, 1)
}
