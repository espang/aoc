package day15

import (
	"fmt"
	"sort"
	"strings"
)

func Part2(input string) {
	xmax := 4_000_000
	ymax := 4_000_000

	sensors := []Sensor{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := parseLine(line)
		dist := distance(sensor, beacon)
		sensors = append(sensors, Sensor{
			Pos:      sensor,
			Distance: dist,
		})
	}
	sort.Slice(sensors, func(i, j int) bool {
		return sensors[i].Distance > sensors[j].Distance
	})

	c := Coordinate{}
	for x := 0; x <= xmax; x++ {
		c.X = x
	yloop:
		for y := 0; y <= ymax; y++ {
			c.Y = y
			for _, sensor := range sensors {
				if dist := distance(c, sensor.Pos); dist <= sensor.Distance {
					// skip over some y values to speed up
					y += (sensor.Distance - dist)
					continue yloop
				}
			}
			fmt.Println(x*4000_000 + y)
			return
		}
	}
	fmt.Println("no solution")
}
