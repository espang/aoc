package day18

import (
	"strconv"
	"strings"
)

var testinput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

type Coord struct {
	X, Y, Z int
}

func ParseLine(line string) Coord {
	vs := []int{}
	for _, number := range strings.Split(line, ",") {
		v, err := strconv.Atoi(number)
		if err != nil {
			panic(err.Error())
		}
		vs = append(vs, v)
	}
	if len(vs) != 3 {
		panic("unexpected number of numbers in given line: " + line)
	}
	return Coord{
		X: vs[0],
		Y: vs[1],
		Z: vs[2],
	}
}

func Parse(input string) []Coord {
	coords := []Coord{}
	for _, line := range strings.Split(input, "\n") {
		coords = append(coords, ParseLine(line))
	}
	return coords
}
