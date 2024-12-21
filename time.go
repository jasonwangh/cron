package cron

import (
	"time"
)

type TimeOffset struct {
	offset time.Duration
}

func NewTimeOffset() *TimeOffset {
	return &TimeOffset{}
}

func (to *TimeOffset) Now() time.Time {
	var now time.Time
	if to.offset != 0 {
		now = time.Now().Add(to.offset)
	} else {
		now = time.Now()
	}
	return now
}

func (to *TimeOffset) GetOffset() time.Duration {
	return to.offset
}

func (to *TimeOffset) SetOffset(sec int64) {
	to.offset += time.Duration(sec) * time.Second
}
