package timeutil

import "time"

type TimeTracker struct {
	timeBegin time.Time
	timeEnd   time.Time
	duration  time.Duration
}

func (t *TimeTracker) Begin() {
	t.timeBegin = time.Now()
}

func (t *TimeTracker) End() {
	t.timeEnd = time.Now()
	t.duration = t.timeEnd.Sub(t.timeBegin)
}

func (t *TimeTracker) GetBeginTimestamp() int64 {
	return t.timeBegin.UTC().Unix()
}

func (t *TimeTracker) GetEndTimestamp() int64 {
	return t.timeEnd.UTC().Unix()
}

func (t *TimeTracker) GetDurationTime() int64 {
	return t.duration.Nanoseconds() / 1000 //us
}
