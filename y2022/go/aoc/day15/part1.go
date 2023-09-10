package day15

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Sensor struct {
	Pos      Coordinate
	Distance int
}

func Part1(input string) {
	y := 2_000_000
	sensors := []Sensor{}
	minX := math.MaxInt
	maxX := math.MinInt
	beacondsAndSensors := CoordinateSet{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := parseLine(line)
		dist := distance(sensor, beacon)
		distFromY := abs(sensor.Y - y)
		if distFromY > dist {
			// sensor is far away from wanted y so we don't care about them
			continue
		}
		sensors = append(sensors, Sensor{Pos: sensor, Distance: dist})
		if x := sensor.X - dist; x < minX {
			minX = x
		}
		if x := sensor.X + dist; x > maxX {
			maxX = x
		}
		beacondsAndSensors.Add(sensor)
		beacondsAndSensors.Add(beacon)
	}
	sort.Slice(sensors, func(i, j int) bool {
		di := sensors[i].Distance - abs(sensors[i].Pos.Y-y)
		dj := sensors[j].Distance - abs(sensors[j].Pos.Y-y)

		return sensors[i].Pos.X-di < sensors[j].Pos.X-dj
	})

	start := 0
	total := 0
	c := Coordinate{X: 0, Y: y}
outer:
	for x := minX; x <= maxX; x++ {
		c.X = x
		if beacondsAndSensors.Contains(c) {
			continue
		}
		for i := start; i < len(sensors); i++ {
			dist := distance(sensors[i].Pos, c)
			if dist <= sensors[i].Distance {
				total++
				continue outer
			}
			// over the sensor, stop checking it
			start = i
		}
	}
	fmt.Println(total)
}
