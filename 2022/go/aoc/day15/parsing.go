package day15

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func Add(c1, c2 Coordinate) Coordinate {
	return Coordinate{
		X: c1.X + c2.X,
		Y: c1.Y + c2.Y,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func distance(c1, c2 Coordinate) int {
	return abs(c1.X-c2.X) + abs(c1.Y-c2.Y)
}

func c(x, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

func into2(s string, on string) (string, string) {
	splitted := strings.Split(s, on)
	if len(splitted) != 2 {
		panic("into2 failed: " + s)
	}
	return splitted[0], splitted[1]
}

// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
func parseLine(s string) (Coordinate, Coordinate) {
	sensorS, beaconS := into2(s, ": ")

	first, second := into2(sensorS, ", ")
	sx, err := strconv.Atoi(strings.TrimPrefix(first, "Sensor at x="))
	if err != nil {
		panic(err)
	}
	sy, err := strconv.Atoi(strings.TrimPrefix(second, "y="))
	if err != nil {
		panic(err)
	}

	first, second = into2(beaconS, ", ")
	bx, err := strconv.Atoi(strings.TrimPrefix(first, "closest beacon is at x="))
	if err != nil {
		panic(err)
	}

	by, err := strconv.Atoi(strings.TrimPrefix(second, "y="))
	if err != nil {
		panic(err)
	}

	return c(sx, sy), c(bx, by)
}

var testinput = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
