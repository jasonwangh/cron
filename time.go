package cron

import (
	"fmt"
	"strconv"
	"time"
)

type TimeOffset struct {
	now    time.Time
	offset time.Duration
}

func NewTimeOffset() *TimeOffset {
	return &TimeOffset{now: time.Now()}
}

func (to *TimeOffset) SetNow() {
	to.now = time.Now().Add(to.offset)
}

func (to *TimeOffset) Now() time.Time {
	return to.now
}

func (to *TimeOffset) Sec() int64 {
	to.SetNow()
	return to.now.Unix()
}

func (to *TimeOffset) Msec() int64 {
	to.SetNow()
	return int64(to.now.UnixNano() / time.Millisecond.Nanoseconds())
}

func (to *TimeOffset) Nsec() int64 {
	to.SetNow()
	return int64(to.now.UnixNano())
}

func (to *TimeOffset) Hour() int {
	to.SetNow()
	return to.now.Hour()
}

func (to *TimeOffset) Minute() int {
	to.SetNow()
	return to.now.Minute()
}

func (to *TimeOffset) Second() int {
	to.SetNow()
	return to.now.Second()
}

func (to *TimeOffset) DayNum() uint32 {
	to.SetNow()
	return uint32(to.now.Unix() / 86400)
}

func (to *TimeOffset) Month() int {
	to.SetNow()
	return int(to.now.Month())
}

func (to *TimeOffset) YearMonthDay(offsetdays ...int) int {
	to.SetNow()
	var offsetday int
	if len(offsetdays) > 0 {
		offsetday = offsetdays[0]
	}
	year, month, day := to.now.AddDate(0, 0, offsetday).Date()
	str := fmt.Sprintf("%04d%02d%02d", year, int(month), day)
	ret, _ := strconv.Atoi(str)
	return ret
}

func (to *TimeOffset) GetOffset() time.Duration {
	return to.offset
}

func (to *TimeOffset) SetOffset(sec int64) {
	to.offset = to.offset + time.Duration(sec)*time.Second
	to.SetNow()
}
