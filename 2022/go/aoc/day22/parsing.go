package day22

import (
	"fmt"
	"strconv"
	"strings"
)

const testinput = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

type Rotate int

const (
	RRight Rotate = iota
	RLeft
)

func (r Rotate) String() string {
	switch r {
	case RRight:
		return "turn right"
	case RLeft:
		return "turn left"
	}
	panic("unreachable")
}

type Move int

func (m Move) String() string {
	return fmt.Sprintf("move %d", m)
}

type Orientation int

const (
	Right Orientation = iota
	Down
	Left
	Up
)

func (o Orientation) String() string {
	switch o {
	case Right:
		return "Right"
	case Left:
		return "Left"
	case Up:
		return "Up"
	case Down:
		return "Down"
	}
	panic("unreachable")
}

func (o Orientation) Rotate(by Rotate) Orientation {
	switch by {
	case RRight:
		switch o {
		case Left:
			return Up
		case Right:
			return Down
		case Up:
			return Right
		case Down:
			return Left
		}
	case RLeft:
		switch o {
		case Left:
			return Down
		case Right:
			return Up
		case Up:
			return Left
		case Down:
			return Right
		}
	}
	panic("unreachable rotation")
}

type Cell int

const (
	Void Cell = iota
	Open
	Wall
)

func Parse(s string) ([][]Cell, []any) {
	splitted := strings.Split(s, "\n\n")
	grid := splitted[0]
	instructions := splitted[1]
	return parseGrid(grid), parseInstruction(instructions)
}

func parseGrid(grid string) Grid {
	matrix := [][]Cell{}
	for i, line := range strings.Split(grid, "\n") {
		row := []Cell{}
		for _, c := range line {
			switch c {
			case ' ':
				row = append(row, Void)
			case '.':
				row = append(row, Open)
			case '#':
				row = append(row, Wall)
			default:
				panic("unexpected char in line " + strconv.Itoa(i) + ". " + line)
			}
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func parseInstruction(instructions string) []any {
	if strings.HasPrefix(instructions, "L") {
		return append([]any{RLeft}, parseInstruction(instructions[1:])...)
	}
	if strings.HasPrefix(instructions, "R") {
		return append([]any{RRight}, parseInstruction(instructions[1:])...)
	}
	numberEndsAt := strings.IndexAny(instructions, "LR")
	if numberEndsAt == -1 {
		v, err := strconv.Atoi(instructions)
		if err != nil {
			panic("lastNumber: " + err.Error())
		}
		return []any{Move(v)}
	}
	v, err := strconv.Atoi(instructions[:numberEndsAt])
	if err != nil {
		panic("number: " + err.Error())
	}
	return append([]any{Move(v)}, parseInstruction(instructions[numberEndsAt:])...)
}
