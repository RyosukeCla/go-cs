package arraypush

import (
	"time"
)

type RateLimit interface {
	CheckLimit() bool
}

type rateLimit struct {
	array    []bool
	duration time.Duration
	limit    int
}

func New(duration time.Duration, limit int) *rateLimit {
	return &rateLimit{
		duration: duration,
		limit:    limit,
		array:    make([]bool, 0, 0),
	}
}

func (rl *rateLimit) CheckLimit() bool {
	current := len(rl.array)
	if current > rl.limit {
		return true
	}
	rl.array = append(rl.array, true)
	time.AfterFunc(rl.duration, func() {
		rl.array = rl.array[1:]
	})

	return false
}
