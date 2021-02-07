package main

import (
	"log"
	"time"
)

type TimestampCounter struct {
	Interval time.Duration
	stamps   []time.Time
	first    int
	repo     TimestampRepo
}

const CompactThreshold = 1

func NewTimestampCounter(interval time.Duration, repo TimestampRepo) (*TimestampCounter, error) {
	stamps, err := repo.LoadAll()
	if err != nil {
		return nil, err
	}
	counter := &TimestampCounter{
		Interval: interval,
		repo:     repo,
		stamps:   stamps,
	}
	return counter, nil
}

func (t *TimestampCounter) AddAndCount(stamp time.Time) int {
	t.add(stamp)
	return t.count(stamp)
}

func (t *TimestampCounter) add(stamp time.Time) {
	t.stamps = append(t.stamps, stamp)

	err := t.repo.AppendOne(stamp)
	if err != nil {
		log.Println(err)
	}
}

func (t *TimestampCounter) count(stamp time.Time) int {
	t.first += t.countOldStamps(stamp)
	count := len(t.stamps) - t.first
	if t.first > 0 && count > 0 && t.first/count > CompactThreshold {
		t.compact()
	}
	return count
}

func (t *TimestampCounter) countOldStamps(stamp time.Time) int {
	intervalStart := stamp.Add(- t.Interval)
	for i, v := range t.stamps[t.first:] { //use binary search?
		if v.After(intervalStart) {
			return i
		}
	}
	return 0
}

func (t *TimestampCounter) compact() {
	t.stamps = t.stamps[t.first:]
	t.first = 0

	err := t.repo.StoreAll(t.stamps)
	if err != nil {
		log.Println(err)
	}
}
