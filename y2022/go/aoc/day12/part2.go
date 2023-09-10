package day12

import (
	"fmt"
	"sort"
)

func Part2(input string) {
	_, end, heights := ParseInput(input)

	steps := []int{}

	for x, row := range heights {
		for y, cell := range row {
			if cell == 1 {
				minSteps := bfs(Point{x, y}, end, heights)
				if minSteps > 0 {
					steps = append(steps, minSteps)
				}
			}
		}
	}
	sort.Ints(steps)
	fmt.Println(steps[0])
}
