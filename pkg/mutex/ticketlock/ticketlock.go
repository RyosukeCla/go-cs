package ticketlock

import (
	"runtime"
	"sync/atomic"
)

type mutex struct {
	nowServing int32
	nextTicket int32
}

func New() mutex {
	return mutex{
		nowServing: int32(0),
		nextTicket: int32(-1),
	}
}

func (m *mutex) Acquire() {
	myTicket := atomic.AddInt32(&m.nextTicket, 1)
	for {
		if myTicket == m.nowServing {
			break
		}
		runtime.Gosched()
	}
}

func (m *mutex) Release() {
	atomic.AddInt32(&m.nowServing, 1)
}
