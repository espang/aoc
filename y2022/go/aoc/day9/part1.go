package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Instruction struct {
	Dir    Direction
	Amount int
}

type Point struct{ x, y int }

type PS map[Point]struct{}

func (ps PS) Add(p Point) PS    { ps[p] = struct{}{}; return ps }
func (ps *PS) Cardinality() int { return len(*ps) }

func parseLine(s string) Instruction {
	val, err := strconv.Atoi(s[2:])
	if err != nil {
		panic("parseLine: " + err.Error())
	}
	switch s[0] {
	case 'R':
		return Instruction{RIGHT, val}
	case 'L':
		return Instruction{LEFT, val}
	case 'D':
		return Instruction{DOWN, val}
	case 'U':
		return Instruction{UP, val}
	}
	panic("parseLine: " + s)
}

func moveHead(head Point, dir Direction) Point {
	switch dir {
	case UP:
		head.y -= 1
	case DOWN:
		head.y += 1
	case LEFT:
		head.x -= 1
	case RIGHT:
		head.x += 1
	}
	return head
}

func moveTail(head, tail Point) Point {
	dx := head.x - tail.x
	dy := head.y - tail.y
	if -1 <= dx && dx <= 1 && -1 <= dy && dy <= 1 {
		return tail
	}
	limit := func(v int) int {
		if v < -1 {
			return -1
		}
		if v > 1 {
			return 1
		}
		return v
	}
	tail.x += limit(dx)
	tail.y += limit(dy)
	return tail
}

func move(head, tail Point, inst Instruction, points PS) (Point, Point, PS) {
	if inst.Amount == 0 {
		return head, tail, points
	}

	head = moveHead(head, inst.Dir)
	tail = moveTail(head, tail)
	points = points.Add(tail)
	inst.Amount--

	return move(head, tail, inst, points)
}

func instructionFromLines(lines []string) []Instruction {
	result := make([]Instruction, 0, len(lines))
	for _, l := range lines {
		result = append(result, parseLine(l))
	}
	return result
}

func Part1(input string) {
	lines := strings.Split(input, "\n")
	instructions := instructionFromLines(lines)
	head := Point{}
	tail := Point{}
	points := PS{}
	for _, inst := range instructions {
		head, tail, points = move(head, tail, inst, points)
	}
	fmt.Println(points.Cardinality())
}
