package day24

import (
	"context"
	"errors"
	"strconv"
	"unicode"

	"github.com/espang/aoc/aoc"
)

type Walls [][]bool

func isWall(walls Walls, loc aoc.Coordinate) bool {
	return walls[loc.Y][loc.X]
}

func potentialMoves(walls Walls, loc aoc.Coordinate) []aoc.Coordinate {
	var moves []aoc.Coordinate
	for _, f := range []aoc.MoveFn{aoc.MoveUp, aoc.MoveDown, aoc.MoveLeft, aoc.MoveRight} {
		pos := f(loc)
		if !isWall(walls, pos) {
			moves = append(moves, pos)
		}
	}
	return moves
}

func parseInput(input string) (Walls, map[int]aoc.Coordinate) {
	var walls Walls
	numbers := map[int]aoc.Coordinate{}
	lines := aoc.SplitBy(input, aoc.Newline)
	for yCord, line := range lines {
		var row []bool
		for xCord, v := range line {
			switch v {
			case '#':
				row = append(row, true)
			case '.':
				row = append(row, false)
			default:
				if unicode.IsDigit(v) {
					numbers[int(v)-48] = aoc.Coordinate{X: xCord, Y: yCord}
					row = append(row, false)
				} else {
					panic("parse input: unexpected rune " + string(v))
				}
			}
		}
		walls = append(walls, row)
	}
	return walls, numbers
}

type state struct {
	where aoc.Coordinate
	// seen will store the numbers this state has seen so far by
	// setting the numbers bit to 1. The example has 4, the input
	// has 7 numbers.
	seen  int
	steps int
}

type stateWithoutSteps struct {
	where aoc.Coordinate
	seen  int
}

func setBit(v int, pos int) int {
	v |= 1 << pos
	return v
}

func checkSeen(numbers []int) func(int) bool {
	search := 0
	for _, k := range numbers {
		search = setBit(search, k)
	}
	return func(i int) bool { return i == search }
}

func part1(content string) (string, error) {
	walls, numbers := parseInput(content)

	toFind := aoc.MapTranspose(numbers)
	isDone := checkSeen(aoc.MapValues(toFind))

	start := numbers[0]

	// visited makes sure to reduce the search space. We are looking for the shortest
	// path visiting all numbers. When we reach the same coordinate, having seen the
	// sambe numbers we can't get a shorter path from there - because of the use of BFS.
	// The first time we reached that setup will be the shortest path to that setup!
	visited := aoc.NewSet(stateWithoutSteps{where: start, seen: setBit(0, 0)})

	queue := aoc.NewQueue(state{where: start, seen: setBit(0, 0)})
	for queue.NotEmpty() {
		current := queue.Pop()

		for _, nextPos := range potentialMoves(walls, current.where) {
			seen := current.seen
			if n, ok := toFind[nextPos]; ok {
				// this is one of the numbers!
				seen = setBit(seen, n)
				if isDone(seen) {
					return strconv.Itoa(current.steps + 1), nil
				}
			}
			sws := stateWithoutSteps{where: nextPos, seen: seen}
			if !visited.Contains(sws) {
				visited.Add(sws)
				queue.Push(state{where: nextPos, seen: seen, steps: current.steps + 1})
			}
		}
	}

	return "", errors.New("no path found")
}

func Part1() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 24)
	content := aoc.LoadFrom(filename)
	return part1(content)
}

func part2(content string) (string, error) {
	walls, numbers := parseInput(content)

	toFind := aoc.MapTranspose(numbers)
	isDone := checkSeen(aoc.MapValues(toFind))

	start := numbers[0]

	// visited makes sure to reduce the search space. We are looking for the shortest
	// path visiting all numbers. When we reach the same coordinate, having seen the
	// sambe numbers we can't get a shorter path from there - because of the use of BFS.
	// The first time we reached that state will be the shortest path to that state!
	visited := aoc.NewSet(stateWithoutSteps{where: start, seen: setBit(0, 0)})

	queue := aoc.NewQueue(state{where: start, seen: setBit(0, 0)})
	for queue.NotEmpty() {
		current := queue.Pop()

		if isDone(current.seen) && current.where.Equal(start) {
			return strconv.Itoa(current.steps), nil
		}

		for _, nextPos := range potentialMoves(walls, current.where) {
			seen := current.seen
			if n, ok := toFind[nextPos]; ok {
				// this is one of the numbers!
				seen = setBit(seen, n)
			}
			sws := stateWithoutSteps{where: nextPos, seen: seen}
			if !visited.Contains(sws) {
				visited.Add(sws)
				queue.Push(state{where: nextPos, seen: seen, steps: current.steps + 1})
			}
		}
	}

	return "", errors.New("no path found")
}

func Part2() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 24)
	content := aoc.LoadFrom(filename)
	return part2(content)
}
