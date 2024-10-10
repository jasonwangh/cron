package cron

import (
	"fmt"
	"strconv"
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

func (to *TimeOffset) Sec() int64 {
	return to.Now().Unix()
}

func (to *TimeOffset) Msec() int64 {
	return int64(to.Now().UnixNano() / time.Millisecond.Nanoseconds())
}

func (to *TimeOffset) Nsec() int64 {
	return int64(to.Now().UnixNano())
}

func (to *TimeOffset) Hour() int {
	return to.Now().Hour()
}

func (to *TimeOffset) Minute() int {
	return to.Now().Minute()
}

func (to *TimeOffset) Second() int {
	return to.Now().Second()
}

func (to *TimeOffset) DayNum() uint32 {
	return uint32(to.Now().Unix() / 86400)
}

func (to *TimeOffset) Month() int {
	return int(to.Now().Month())
}

func (to *TimeOffset) YearMonthDay(offsetdays ...int) int {
	var offsetday int
	if len(offsetdays) > 0 {
		offsetday = offsetdays[0]
	}
	year, month, day := to.Now().AddDate(0, 0, offsetday).Date()
	str := fmt.Sprintf("%04d%02d%02d", year, int(month), day)
	ret, _ := strconv.Atoi(str)
	return ret
}

func (to *TimeOffset) GetOffset() time.Duration {
	return to.offset
}

func (to *TimeOffset) SetOffset(sec int64) {
	to.offset = to.offset + time.Duration(sec)*time.Second
}
