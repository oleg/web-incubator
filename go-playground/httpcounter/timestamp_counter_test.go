package main

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewTimestampCounter_loads_stamps_from_repo(t *testing.T) {
	stamps := []time.Time{time.Now(), time.Now(), time.Now()}
	repo := &testTimestampsRepo{stamps: stamps}

	counter, err := NewTimestampCounter(time.Duration(7)*time.Second, repo)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(counter.stamps, stamps) {
		t.Errorf("counter stamps %v are not equal to expected %v", counter.stamps, stamps)
	}
}

func TestNewTimestampCounter_returns_error_if_repo_returns_error(t *testing.T) {
	someError := errors.New("some error")
	repo := &testTimestampsRepo{err: someError}

	_, err := NewTimestampCounter(time.Duration(9)*time.Second, repo)

	if err != someError {
		t.Errorf("Incorrect error %v returned, expected %v", err, someError)
	}
}

func TestTimestampCounter_Add_adds_time_to_the_stamps(t *testing.T) {
	ts := TimestampCounter{
		repo: &testTimestampsRepo{},
	}
	now := time.Now()

	ts.add(now)

	if len(ts.stamps) != 1 || ts.stamps[0] != now {
		t.Errorf("add should add time %v to stamps %v", now, ts.stamps)
	}
}

func TestTimestampCounter_Count_counts_timestamps_that_are_in_interval(t *testing.T) {
	data := timestampsEachSeconds(10)
	first := data[0]

	tests := []struct {
		name          string
		stamps        []time.Time
		from          time.Time
		expectedCount int
	}{
		{"1 sec", data[:1], data[0], 1},
		{"2 sec", data[:2], data[1], 2},
		{"3 sec", data[:3], data[2], 3},
		{"4 sec", data[:4], data[3], 4},
		{"5 sec", data[:5], data[4], 5},
		{"6 sec", data[:6], data[5], 5},
		{"7 sec", data[:7], data[6], 5},
		{"8 sec", data[:8], data[7], 5},
		{"9 sec", data[:9], data[8], 5},
		{"10 sec", data, data[9], 5},
		{"11 sec", data, first.Add(time.Duration(10) * time.Second), 4},
		{"12 sec", data, first.Add(time.Duration(11) * time.Second), 3},
		{"13 sec", data, first.Add(time.Duration(12) * time.Second), 2},
		{"14 sec", data, first.Add(time.Duration(13) * time.Second), 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := TimestampCounter{
				Interval: 5 * time.Second,
				stamps:   test.stamps,
				repo:     &testTimestampsRepo{},
			}

			count := ts.count(test.from)

			if count != test.expectedCount {
				t.Errorf("Count for %v is %v, expected %v", test.from, count, test.expectedCount)
			}
		})
	}

}

func TestTimestampCounter_AddAndCount_ignores_older_entries(t *testing.T) {
	data := timestampsEachSeconds(10)
	first := data[0]

	ts := TimestampCounter{
		Interval: 5 * time.Second,
		stamps:   data,
		repo:     &testTimestampsRepo{},
	}

	count := ts.AddAndCount(first.Add(time.Duration(20) * time.Second))
	if count != 1 {
		t.Errorf("Wrong count %v, epected %v", count, 1)
	}
	count = ts.AddAndCount(first.Add(time.Duration(21) * time.Second))
	if count != 2 {
		t.Errorf("Wrong count %v epected %v", count, 2)
	}
	count = ts.AddAndCount(first.Add(time.Duration(22) * time.Second))
	if count != 3 {
		t.Errorf("Wrong count %v epected %v", count, 3)
	}
	count = ts.AddAndCount(first.Add(time.Duration(23) * time.Second))
	if count != 4 {
		t.Errorf("Wrong count %v epected %v", count, 4)
	}
	count = ts.AddAndCount(first.Add(time.Duration(24) * time.Second))
	if count != 5 {
		t.Errorf("Wrong count %v epected %v", count, 5)
	}
}

func TestTimestamps_Count_cleanups_the_slice(t *testing.T) {
	data := timestampsEachSeconds(10)

	ts := TimestampCounter{
		Interval: 2 * time.Second,
		stamps:   data,
		repo:     &testTimestampsRepo{},
	}

	_ = ts.count(data[9])

	if len(ts.stamps) != 2 {
		t.Errorf("Wrong length %v, expected %v", len(ts.stamps), 2)
	}
}

func timestampsEachSeconds(n int) []time.Time {
	now := time.Now()
	stamps := make([]time.Time, n)
	for i := 0; i < n; i++ {
		stamps[i] = now.Add(time.Duration(i) * time.Second)
	}
	return stamps
}
