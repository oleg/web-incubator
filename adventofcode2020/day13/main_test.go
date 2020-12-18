package main

import (
	"testing"
)

func Test_day13_task1_parse(t *testing.T) {
	time, ids := parseSchedule("939\n7,13,x,x,59,x,31,19")

	if time != 939 {
		t.Errorf("Wrong time %v", time)
	}
	if len(ids) != 5 || ids[0] != 7 || ids[4] != 19 {
		t.Errorf("Wrong ids %v", ids)
	}
}

func Test_day13_task1_findBus(t *testing.T) {
	time, ids := parseSchedule("939\n7,13,x,x,59,x,31,19")

	b := findBus(time, ids)

	if b.id != 59 || b.waitTime != 5 || b.departure != 944 {
		t.Errorf("Wrong b info %v", b)
	}
}

func Test_day13_task2_parseBuses(t *testing.T) {
	buses := parseBuses("939\n7,13,x,x,59,x,31,19")

	//7,13,x,x,59,x,31,19
	if len(buses) != 5 {
		t.Errorf("Wrong length %v", len(buses))
	}
	if buses[2].id != 59 || buses[2].offset != 3 {
		t.Errorf("Wrong buses[2] value %v", buses[2])
	}
	if buses[4].id != 19 || buses[4].offset != 1 {
		t.Errorf("Wrong buses[2] value %v", buses[4])
	}
}

func Test_day13_task2_find(t *testing.T) {
	tests := []struct {
		schedule string
		time     int
	}{
		{"\n17,x,13,19", 3417},
		{"\n67,7,59,61", 754018},
		{"\n67,x,7,59,61", 779210},
		{"\n67,7,x,59,61", 1261476},
		{"\n7,13,x,x,59,x,31,19", 1068781},
		{"\n1789,37,47,1889", 1202161486},
		//19*53272200849006
	}
	for _, test := range tests {
		t.Run(test.schedule, func(t *testing.T) {

			time := findTime(parseBuses(test.schedule))
			if time != test.time {
				t.Errorf("Wrong time %v expected %v", time, test.time)
			}
		})
	}
}
