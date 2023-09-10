package day17

import (
	"errors"
	"slices"
	"strconv"

	"github.com/espang/aoc/aoc"
)

func Part1() (string, error) {
	passcode := "awrkjxxr"

	queue := aoc.NewQueue[[]byte](nil)
	for queue.NotEmpty() {
		path := queue.Pop()
		pos := followPath(path, aoc.Coordinate{})

		for _, direction := range OpenDoors(passcode, path) {
			nextPos := aoc.Move(direction)(pos)
			if nextPos.Equal(aoc.Coordinate{X: 3, Y: -3}) {
				return string(append(path, directionToPathElement(direction))), nil
			}

			if insideGrid(nextPos) {
				path = slices.Clone(path)
				queue.Push(append(path, directionToPathElement(direction)))
			}
		}
	}
	return "", errors.New("no path found")
}

func Part2() (string, error) {
	passcode := "awrkjxxr"
	longestPath := 0

	queue := aoc.NewQueue[[]byte](nil)
	for queue.NotEmpty() {
		path := queue.Pop()
		pos := followPath(path, aoc.Coordinate{})

		for _, direction := range OpenDoors(passcode, path) {
			nextPos := aoc.Move(direction)(pos)

			if nextPos.Equal(aoc.Coordinate{X: 3, Y: -3}) {
				pathLength := len(path) + 1
				if pathLength > longestPath {
					longestPath = pathLength
				}
			} else if insideGrid(nextPos) {
				path = slices.Clone(path)
				queue.Push(append(path, directionToPathElement(direction)))
			}
		}
	}

	if longestPath > 0 {
		return strconv.Itoa(longestPath), nil
	}
	return "", errors.New("no path found")
}
