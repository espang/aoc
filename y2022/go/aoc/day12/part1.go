package day12

import (
	"container/list"
	"fmt"
	"math"
)

func heightOf(p Point, heights [][]int) int {
	if 0 <= p.x && p.x < len(heights) {
		row := heights[p.x]
		if 0 <= p.y && p.y < len(row) {
			return row[p.y]
		}
		return math.MaxInt
	}
	return math.MaxInt
}

func plus(p1, p2 Point) Point {
	return Point{
		x: p1.x + p2.x,
		y: p1.y + p2.y,
	}
}

func equal(p1, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func potentialSteps(pos Point, heights [][]int) (neighbours []Point) {
	deltas := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	reachableHeight := heightOf(pos, heights) + 1

	for _, delta := range deltas {
		newPos := plus(pos, delta)
		h := heightOf(newPos, heights)
		if h <= reachableHeight {
			neighbours = append(neighbours, newPos)
		}
	}
	return neighbours
}

func bfs(start, end Point, heights [][]int) int {
	stepsTo := map[Point]int{
		start: 0,
	}
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Point)
		steps := stepsTo[current]

		if equal(current, end) {
			return steps
		}

		neighbours := potentialSteps(current, heights)
		for _, nextPosition := range neighbours {
			if _, ok := stepsTo[nextPosition]; ok {
				continue
			}
			stepsTo[nextPosition] = steps + 1
			queue.PushBack(nextPosition)
		}
	}
	return -1
}

func Part1(input string) {
	start, end, heights := ParseInput(input)
	fmt.Println(bfs(start, end, heights))
}
