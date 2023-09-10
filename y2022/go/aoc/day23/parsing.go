package day23

import "strings"

var testinput = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func Parse(input string) []Coord {
	var coords []Coord
	for y, line := range strings.Split(input, "\n") {
		for x, cell := range line {
			switch cell {
			case '.':
			case '#':
				coords = append(coords, Coord{x: x, y: -y})
			}
		}
	}
	return coords
}
