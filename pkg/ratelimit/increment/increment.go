package increment

import (
	"time"
)

type RateLimit interface {
	CheckLimit() bool
}

type rateLimit struct {
	counter  int
	duration time.Duration
	limit    int
}

func New(duration time.Duration, limit int) *rateLimit {
	return &rateLimit{
		duration: duration,
		limit:    limit,
		counter:  0,
	}
}

func (rl *rateLimit) CheckLimit() bool {
	current := rl.counter
	if current > rl.limit {
		return true
	}
	rl.counter = rl.counter + 1
	time.AfterFunc(rl.duration, func() {
		rl.counter = rl.counter - 1
	})

	return false
}
