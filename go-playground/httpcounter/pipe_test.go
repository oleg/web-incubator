package main

import (
	"reflect"
	"testing"
	"time"
)

func TestMakePipe(t *testing.T) {
	t1 := time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	t2 := time.Date(2020, 2, 2, 2, 2, 2, 2, time.UTC)
	t3 := time.Date(2020, 3, 3, 3, 3, 3, 3, time.UTC)

	times := make([]time.Time, 0)
	seconds := make([]int, 0)
	in, out, pipe := makePipe(func(stamp time.Time) int {
		times = append(times, stamp)
		return stamp.Second()
	})

	go pipe()
	in <- t1
	seconds = append(seconds, <-out)
	in <- t2
	seconds = append(seconds, <-out)
	in <- t3
	seconds = append(seconds, <-out)

	expectedTimes := []time.Time{t1, t2, t3}
	if !reflect.DeepEqual(times, expectedTimes) {
		t.Errorf("Wrong times %v, expected %v", times, expectedTimes)
	}
	expectedSeconds := []int{1, 2, 3}
	if !reflect.DeepEqual(seconds, expectedSeconds) {
		t.Errorf("Wrong seconds %v, expected %v", seconds, expectedSeconds)
	}
}

