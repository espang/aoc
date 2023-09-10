package day1

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/espang/aoc/aoc"
)

type Tuple struct {
	Dir   rune
	Steps int
}

func parseInput(input string) ([]Tuple, error) {
	var allErrors error
	f := func(element string) Tuple {
		l, n, err := aoc.ParseLetterNumber(element)
		if err != nil {
			allErrors = errors.Join(allErrors, err)
			return Tuple{}
		}
		if l != 'L' && l != 'R' {
			allErrors = errors.Join(allErrors, fmt.Errorf("unexpected letter: %v", l))
			return Tuple{}
		}
		return Tuple{
			Dir:   l,
			Steps: n,
		}
	}
	entries := aoc.SplitBy(input, aoc.CommaSpace)
	tuples := aoc.SliceMap(entries, f)
	return tuples, allErrors
}

func Part1() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 1)
	content := aoc.LoadFrom(filename)
	tuples, err := parseInput(content)
	if err != nil {
		return "", err
	}

	location := part1(tuples)
	distanceFromStart := aoc.ManhattanDistance(aoc.Coordinate{}, location)
	return strconv.Itoa(distanceFromStart), nil
}

func part1(tuples []Tuple) aoc.Coordinate {
	location := aoc.Coordinate{}
	direction := aoc.Up
	for _, t := range tuples {
		switch t.Dir {
		case 'L':
			direction = direction.Left()
		case 'R':
			direction = direction.Right()
		default:
			panic("unreachable")
		}

		switch direction {
		case aoc.Up:
			location.Y += t.Steps
		case aoc.Down:
			location.Y -= t.Steps
		case aoc.Left:
			location.X -= t.Steps
		case aoc.Right:
			location.X += t.Steps
		default:
			panic("unreachable")
		}
	}
	return location
}

func Part2() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 1)
	content := aoc.LoadFrom(filename)
	tuples, err := parseInput(content)
	if err != nil {
		return "", err
	}

	location, found := part2(tuples)
	if !found {
		return "", errors.New("no suitable location found")
	}

	distanceFromStart := aoc.ManhattanDistance(aoc.Coordinate{}, location)
	return strconv.Itoa(distanceFromStart), nil
}

func part2(tuples []Tuple) (aoc.Coordinate, bool) {
	location := aoc.Coordinate{}
	visited := aoc.NewSet[aoc.Coordinate]()
	visited.Add(location)
	direction := aoc.Up
	for _, t := range tuples {
		switch t.Dir {
		case 'L':
			direction = direction.Left()
		case 'R':
			direction = direction.Right()
		default:
			panic("unreachable")
		}

		move := aoc.Move(direction)
		for i := 0; i < t.Steps; i++ {
			location = move(location)
			if visited.Contains(location) {
				return location, true
			}
			visited.Add(location)
		}
	}
	return location, false
}
