package main

import (
	"time"
)

type Timestamps struct {
	Interval         time.Duration
	CleanupThreshold int //should be configured according to specific load
	stamps           []time.Time
	first            int
}

func (t *Timestamps) AddAndCount(stamp time.Time) int {
	t.Add(stamp)
	return t.Count(stamp)
}

func (t *Timestamps) Add(stamp time.Time) {
	t.stamps = append(t.stamps, stamp)
}

func (t *Timestamps) Count(stamp time.Time) int { //counts changes the state of t.first
	intervalStart := stamp.Add(- t.Interval)
	for i, v := range t.stamps[t.first:] {
		//fmt.Printf("i   :%v\nfrom:%v\nv   :%v\naftr:%v\n\n", i, intervalStart, v, v.After(intervalStart))
		if v.After(intervalStart) { //not before
			t.first = i
			break
		}

	}
	count := len(t.stamps) - t.first
	//if len(t.stamps)%1_000 == 0 {
	//	fmt.Printf("%v, %v, %v\n", t.first, len(t.stamps), counts)
	//}
	if t.first > t.CleanupThreshold {
		//fmt.Printf("%v\n", "HOUSKEEPING")
		t.stamps = t.stamps[t.first:]
		t.first = 0
	}
	return count
}
