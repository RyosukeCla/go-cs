package spinlock

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	lock uint32
}

func New() Mutex {
	return Mutex{
		lock: uint32(0),
	}
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapUint32(&m.lock, 0, 1) {
			break
		}
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	atomic.StoreUint32(&m.lock, 0)
}
