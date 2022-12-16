package day15

import (
	"fmt"
	"strings"
)

func Part2(input string) {
	xmax := 4_000_000
	ymax := 4_000_000

	distancesBySensor := map[Coordinate]int{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := parseLine(line)
		distancesBySensor[sensor] = distance(sensor, beacon)
	}
	c := Coordinate{}
	for x := 0; x <= xmax; x++ {
		c.X = x
	yloop:
		for y := 0; y <= ymax; y++ {
			c.Y = y
			for sc, d := range distancesBySensor {
				if dist := distance(c, sc); dist <= d {
					// skip over some y values to speed up
					y += (d - dist)
					continue yloop
				}
			}
			fmt.Println(x*4000_000 + y)
			return
		}
	}
	fmt.Println("no solution")
}
