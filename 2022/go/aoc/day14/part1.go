package day14

import "fmt"

var testinput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

type PointSet map[Point]struct{}

func (ps *PointSet) Add(p Point) { (*ps)[p] = struct{}{} }
func (ps *PointSet) Contains(p Point) bool {
	_, ok := (*ps)[p]
	return ok
}
func (ps *PointSet) Union(ps2 PointSet) {
	for k := range ps2 {
		ps.Add(k)
	}
}
func (ps *PointSet) Cardinality() int { return len(*ps) }

func direction(from, to Point) Point {
	limit := func(i int) int {
		if i < -1 {
			return -1
		}
		if i > 1 {
			return 1
		}
		return i
	}
	dx := to.x - from.x
	dy := to.y - from.y
	return p(limit(dx), limit(dy))
}

func toRocks(points []Point) PointSet {
	if len(points) == 0 {
		return PointSet{}
	}
	ps := PointSet{}
	currentPosition := points[0]
	for _, point := range points[1:] {
		step := direction(currentPosition, point)
		for currentPosition != point {
			ps.Add(currentPosition)
			currentPosition = Add(currentPosition, step)
		}
		ps.Add(currentPosition)
	}
	return ps
}

func dropSand(at Point, maxY int, rocks *PointSet) bool {
	if at.y > maxY {
		return false
	}

	// fall down
	down := Add(at, p(0, 1))
	if !rocks.Contains(down) {
		return dropSand(down, maxY, rocks)
	}
	left := Add(at, p(-1, 1))
	if !rocks.Contains(left) {
		return dropSand(left, maxY, rocks)
	}

	right := Add(at, p(1, 1))
	if !rocks.Contains(right) {
		return dropSand(right, maxY, rocks)
	}

	// sand doesn't move, so it is stable
	// stable sand becomes rock
	rocks.Add(at)
	return true
}

func Part1(input string) {
	ppoints := parseInput(input)
	// find maxY to detect non stable sand
	maxY := 0
	rocks := PointSet{}
	for _, points := range ppoints {
		rocks.Union(toRocks(points))
		for _, point := range points {
			if point.y > maxY {
				maxY = point.y
			}
		}

	}
	before := rocks.Cardinality()

	for dropSand(p(500, 0), maxY, &rocks) {
	}

	fmt.Println(rocks.Cardinality() - before)
}
