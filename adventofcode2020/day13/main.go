package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strconv"
	"strings"
)

func main() {
	schedule := misc.MustReadFileToString("day13/input.txt")

	bus := findBus(parseSchedule(schedule))
	println(bus.id * bus.waitTime)

	time := findTime(parseBuses(schedule))
	println(time)
}

type bus struct {
	id, offset int
}

type busInfo struct {
	id, departure, waitTime int
}

func findBus(time int, ids []int) busInfo {
	buses := make([]busInfo, 0, len(ids))

	for _, id := range ids {
		buses = append(buses, findEarliest(id, time))
	}
	minBus := buses[0]
	for _, b := range buses[1:] {
		if b.waitTime < minBus.waitTime {
			minBus = b
		}
	}
	return minBus
}

func findEarliest(id, time int) busInfo {
	end := id
	for end < time {
		end += id
	}
	return busInfo{id: id, departure: end, waitTime: end - time}
}

func parseSchedule(schedule string) (int, []int) {
	timeAndIds := strings.Split(schedule, "\n")
	time := misc.MustAtoi(timeAndIds[0])

	ids := make([]int, 0)
	for _, idStr := range strings.Split(timeAndIds[1], ",") {
		if num, err := strconv.Atoi(idStr); err == nil {
			ids = append(ids, num)
		}
	}
	return time, ids
}

func parseBuses(schedule string) []bus {
	timeAndIds := strings.Split(schedule, "\n")
	ids := strings.Split(timeAndIds[1], ",")

	buses := make([]bus, 0)
	offset := 0
	for _, idStr := range ids {
		num, err := strconv.Atoi(idStr)
		if err != nil {
			offset++
		} else {
			buses = append(buses, bus{id: num, offset: offset})
			offset = 1
		}
	}
	return buses
}

//todo solve differently
func findTime(buses []bus) int {
	times := make([]int, len(buses))
	for i, b := range buses {
		times[i] = b.id
	}
	for !isCorrect(times, buses) {
		for i := 1; i < len(times); i++ {
			for {
				for times[i-1] >= times[i] {
					times[i] += buses[i].id
				}
				if times[i-1] != times[i]-buses[i].offset {
					times[i-1] += buses[i-1].id
				} else {
					break
				}
			}
		}

	}
	return times[0]
}

func isCorrect(times []int, buses []bus) bool {
	for i := 1; i < len(times); i++ {
		if times[i-1] != times[i]-buses[i].offset {
			return false
		}
	}
	return true
}
