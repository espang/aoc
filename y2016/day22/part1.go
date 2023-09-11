package day22

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/espang/aoc/aoc"
)

type Node struct {
	id        int
	xy        aoc.Coordinate
	size      int
	used      int
	available int
	usedRatio int
}

type CanMove struct {
	FromID, ToID int
}

func From(id int, line string) (Node, error) {
	splitted := strings.Split(line, " ")
	splitted = slices.DeleteFunc(splitted, func(s string) bool { return s == "" })
	if len(splitted) != 5 {
		return Node{}, fmt.Errorf("unexpected line: %s", line)
	}
	x, y, err := parseName(splitted[0])
	size, sErr := toNumber(splitted[1])
	used, uErr := toNumber(splitted[2])
	avai, aErr := toNumber(splitted[3])
	usedR, u2Err := toNumber(splitted[4])
	return Node{
		id:        id,
		xy:        aoc.Coordinate{X: x, Y: y},
		size:      size,
		used:      used,
		available: avai,
		usedRatio: usedR,
	}, errors.Join(err, sErr, uErr, aErr, u2Err)
}

func parseName(s string) (int, int, error) {
	splitted := strings.Split(s, "-")
	if len(splitted) != 3 {
		return 0, 0, fmt.Errorf("invalid name %s", s)
	}
	x, xErr := strconv.Atoi(splitted[1][1:])
	y, yErr := strconv.Atoi(splitted[2][1:])
	if err := errors.Join(xErr, yErr); err != nil {
		return 0, 0, fmt.Errorf("line: %s: error: %v", s, err)
	}
	return x, y, nil
}

func toNumber(s string) (int, error) {
	return strconv.Atoi(s[:len(s)-1])
}

func parseInput(content string) ([]Node, error) {
	lines := aoc.SplitBy(content, aoc.Newline)
	lines = aoc.SliceDrop(2, lines)
	var allErrors error
	id := 0
	f := func(s string) Node {
		n, err := From(id, s)
		if err != nil {
			allErrors = errors.Join(err, allErrors)
		}
		id++
		return n
	}
	return aoc.Map(f, lines), allErrors
}

// viablePair returns wether the data from n1 could move to n2.
func viablePair(n1, n2 Node) bool {
	if n1.id == n2.id {
		return false
	}
	return (n1.used != 0) && (n1.used <= n2.available)
}

func allViablePairs(nodes []Node) []CanMove {
	moves := []CanMove{}
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if viablePair(n1, n2) {
				moves = append(moves, CanMove{FromID: n1.id, ToID: n2.id})
			}
		}
	}
	return moves
}

func Part1() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 22)
	content := aoc.LoadFrom(filename)

	nodes, err := parseInput(content)
	if err != nil {
		return "", err
	}
	moves := allViablePairs(nodes)
	return strconv.Itoa(len(moves)), nil
}

func part2(content string) (string, error) {
	nodes, err := parseInput(content)
	if err != nil {
		return "", err
	}

	last := 0
	for _, n := range nodes {
		if n.xy.X != last {
			last = n.xy.X
			fmt.Println()
		}
		switch {
		case n.used == 0:
			fmt.Print("_")
		case n.used >= 100:
			fmt.Print("#")
		default:
			fmt.Print(".")
		}
	}
	return "", nil
}

func Part2() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 22)
	content := aoc.LoadFrom(filename)
	return part2(content)
}
