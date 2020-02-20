package spinlock

import (
	"runtime"
	"sync/atomic"
)

type mutex struct {
	lock uint32
}

func New() mutex {
	return mutex{
		lock: uint32(0),
	}
}

func (m *mutex) Lock() {
	for {
		if atomic.CompareAndSwapUint32(&m.lock, 0, 1) {
			break
		}
		runtime.Gosched()
	}
}

func (m *mutex) Unlock() {
	atomic.StoreUint32(&m.lock, 0)
}
