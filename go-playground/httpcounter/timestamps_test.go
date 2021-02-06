package main

import (
	"testing"
	"time"
)

func TestTimestamps_Add_adds_time_to_the_stamps(t *testing.T) {
	ts := Timestamps{}
	now := time.Now()

	ts.Add(now)

	if len(ts.stamps) != 1 || ts.stamps[0] != now {
		t.Errorf("add should add time %v to stamps %v", now, ts.stamps)
	}
}

func TestTimestamps_Count_counts_timestamps_that_are_in_interval(t *testing.T) {
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
			ts := Timestamps{
				Interval:         5 * time.Second,
				CleanupThreshold: 1_000,
				stamps:           test.stamps,
			}

			count := ts.Count(test.from)

			if count != test.expectedCount {
				t.Errorf("Count for %v is %v, expected %v", test.from, count, test.expectedCount)
			}
		})
	}

}
func TestTimestamps_Count_cleanups_the_slice(t *testing.T) {
	data := timestampsEachSeconds(10)

	ts := Timestamps{
		Interval:         2 * time.Second,
		CleanupThreshold: 5,
		stamps:           data,
	}

	_ = ts.Count(data[9])

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
