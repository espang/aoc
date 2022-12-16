package day15

import (
	"container/list"
	"fmt"
	"strings"
)

type CoordinateSet map[Coordinate]struct{}

func (cs *CoordinateSet) Add(c Coordinate) {
	(*cs)[c] = struct{}{}
}

func (cs CoordinateSet) Contains(c Coordinate) bool {
	_, ok := cs[c]
	return ok
}

func (cs CoordinateSet) Difference(cs2 CoordinateSet) CoordinateSet {
	result := CoordinateSet{}
	for k := range cs {
		if !cs2.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func (cs CoordinateSet) Union(cs2 CoordinateSet) CoordinateSet {
	result := CoordinateSet{}
	for k := range cs {
		result.Add(k)
	}
	for k := range cs2 {
		result.Add(k)
	}
	return result
}

// given a sensor and a beacon, returns all points that
// can't have a beacon
func handleSensor(sensor Coordinate, y, maxDinstance int) CoordinateSet {
	directions := []Coordinate{{-1, 0}, {1, 0}}
	result := CoordinateSet{}
	queue := list.New()
	start := Coordinate{X: sensor.X, Y: y}
	if distance(sensor, start) > maxDinstance {
		return result
	}
	queue.PushBack(start)
	result.Add(start)

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Coordinate)
		for _, dir := range directions {
			newPos := Add(current, dir)
			if result.Contains(newPos) {
				continue
			}

			if distance(sensor, newPos) <= maxDinstance {
				result.Add(newPos)
				queue.PushBack(newPos)
			}
		}
	}
	return result
}

func Part1(input string) {
	y := 2000000

	allSensors := CoordinateSet{}
	allBeacons := CoordinateSet{}
	allFree := CoordinateSet{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := parseLine(line)
		allBeacons.Add(beacon)
		allSensors.Add(sensor)
		free := handleSensor(sensor, y, distance(sensor, beacon))
		allFree = allFree.Union(free)
	}
	total := 0

	for k := range allFree {
		if allBeacons.Contains(k) || allSensors.Contains(k) {
			continue
		}
		if k.Y == y {
			total++
		}
	}
	fmt.Println(total)
}
