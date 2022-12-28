package day23

import (
	"aoc/aoc"
	"fmt"
)

func Part2(input string) {
	coords := Parse(input)
	s := aoc.Set[Coord]{}
	for _, c := range coords {
		s.Add(c)
	}
	var elvesMoved int
	for turn := 0; ; turn++ {
		s, elvesMoved = step(turn, s)
		if elvesMoved == 0 {
			fmt.Println(turn + 1)
			return
		}
	}
}
