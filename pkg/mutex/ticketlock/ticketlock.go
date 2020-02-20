package ticketlock

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	nowServing int32
	nextTicket int32
}

func New() Mutex {
	return Mutex{
		nowServing: int32(0),
		nextTicket: int32(-1),
	}
}

func (m *Mutex) Acquire() {
	myTicket := atomic.AddInt32(&m.nextTicket, 1)
	for {
		if myTicket == m.nowServing {
			break
		}
		runtime.Gosched()
	}
}

func (m *Mutex) Release() {
	atomic.AddInt32(&m.nowServing, 1)
}
