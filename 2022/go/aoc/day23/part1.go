package day23

import (
	"aoc/aoc"
	"fmt"
	"math"
)

type Direction int

const (
	N Direction = iota
	S
	W
	E
)

type Coord struct {
	x, y int
}

var coordsNeighbours = []Coord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var coordsN = []Coord{{-1, 1}, {0, 1}, {1, 1}}
var coordsS = []Coord{{-1, -1}, {0, -1}, {1, -1}}
var coordsW = []Coord{{-1, -1}, {-1, 0}, {-1, 1}}
var coordsE = []Coord{{1, -1}, {1, 0}, {1, 1}}

func elvesIn(at Coord, coords []Coord, elves aoc.Set[Coord]) int {
	total := 0
	for _, c := range coords {
		c.x += at.x
		c.y += at.y
		if elves.IsMember(c) {
			total++
		}
	}
	return total
}

func moveElf(at Coord, turn int, elves aoc.Set[Coord]) Coord {
	if neighbours := elvesIn(at, coordsNeighbours, elves); neighbours == 0 {
		return at
	}
	for i := turn; i < turn+4; i++ {
		direction := Direction(i % 4)
		var coords []Coord
		var move Coord
		switch direction {
		case N:
			coords = coordsN
			move = Coord{0, 1}
		case S:
			coords = coordsS
			move = Coord{0, -1}
		case W:
			coords = coordsW
			move = Coord{-1, 0}
		case E:
			coords = coordsE
			move = Coord{1, 0}
		}
		if elvesInDirection := elvesIn(at, coords, elves); elvesInDirection == 0 {
			move.x += at.x
			move.y += at.y
			return move
		}
	}
	return at
}

func step(turn int, elves aoc.Set[Coord]) (aoc.Set[Coord], int) {
	proposedStepsTo := map[Coord]int{}
	proposedSteps := map[Coord]Coord{}

	// propose steps
	for elf := range elves {
		proposedStep := moveElf(elf, turn, elves)
		proposedStepsTo[proposedStep]++
		proposedSteps[elf] = proposedStep
	}
	// move
	newPositions := aoc.Set[Coord]{}
	totalMoves := 0
	for currentPosition, nextPosition := range proposedSteps {
		if proposedStepsTo[nextPosition] > 1 {
			newPositions.Add(currentPosition)
		} else {
			if nextPosition != currentPosition {
				totalMoves++
			}
			newPositions.Add(nextPosition)
		}
	}
	return newPositions, totalMoves
}

func findMinMax(elves aoc.Set[Coord]) (Coord, Coord) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for elf := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	return Coord{x: minX, y: minY}, Coord{x: maxX, y: maxY}
}

func show(elves aoc.Set[Coord]) {
	min, max := findMinMax(elves)
	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			if elves.IsMember(Coord{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Part1(input string) {
	coords := Parse(input)
	s := aoc.Set[Coord]{}
	for _, c := range coords {
		s.Add(c)
	}
	for turn := 0; turn < 10; turn++ {
		s, _ = step(turn, s)
	}
	min, max := findMinMax(s)

	positions := (max.x - min.x + 1) * (max.y - min.y + 1)
	fmt.Println(positions - len(s))
}
