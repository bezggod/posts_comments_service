package pkg

import "time"

type Timer struct{}

func NewTimer() *Timer {
	return &Timer{}
}

func (t Timer) NowMoscow() time.Time {
	return time.Now().UTC().Add(3 * time.Hour)
}
