package day9

import (
	"fmt"
	"strings"
)

func moveKnots(knots []Point, inst Instruction, points PS) ([]Point, PS) {
	if inst.Amount == 0 {
		return knots, points
	}

	last := moveHead(knots[0], inst.Dir)
	newKnots := []Point{last}
	for i := 1; i < len(knots); i++ {
		last = moveTail(last, knots[i])
		newKnots = append(newKnots, last)
	}
	points = points.Add(last)
	inst.Amount--

	return moveKnots(newKnots, inst, points)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")
	instructions := instructionFromLines(lines)
	knots := make([]Point, 10)
	points := PS{}
	for _, inst := range instructions {
		knots, points = moveKnots(knots, inst, points)
	}
	fmt.Println(points.Cardinality())
}
