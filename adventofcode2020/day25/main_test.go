package main

import (
	"testing"
)

func Test_day25_loopSize(t *testing.T) {
	e := encryption{
		subjectNumber: 7,
		cardPublicKey: 5764801,
		doorPublicKey: 17807724,
	}

	e.findLoopSizes()
	if e.cardLoopSize != 8 {
		t.Errorf("Wrong card loop size %v", e.cardLoopSize)
	}
	if e.doorLoopSize != 11 {
		t.Errorf("Wrong door loop size %v", e.doorLoopSize)
	}
}
func Test_day25_key(t *testing.T) {
	e := encryption{
		subjectNumber: 7,
		cardPublicKey: 5764801,
		doorPublicKey: 17807724,
	}

	e.findLoopSizes()
	key := e.key()
	if key != 14897079 {
		t.Errorf("Wrong key %v", key)
	}
}
