package day14

import "fmt"

func atBottom(maxY int, p Point) bool {
	return p.y == maxY+1
}

// maxY has to be the y value above the infinite layer of rocks
func dropSand2(at Point, maxY int, rocks *PointSet) {
	if atBottom(maxY, at) {
		rocks.Add(at)
		return
	}
	// fall down
	down := Add(at, p(0, 1))
	if !rocks.Contains(down) {
		dropSand2(down, maxY, rocks)
		return
	}
	left := Add(at, p(-1, 1))
	if !rocks.Contains(left) {
		dropSand2(left, maxY, rocks)
		return
	}

	right := Add(at, p(1, 1))
	if !rocks.Contains(right) {
		dropSand2(right, maxY, rocks)
		return
	}

	// sand doesn't move, so it is stable
	// stable sand becomes rock
	rocks.Add(at)
}

func Part2(input string) {
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

	start := p(500, 0)
	for !rocks.Contains(start) {
		dropSand2(p(500, 0), maxY, &rocks)
	}

	fmt.Println(rocks.Cardinality() - before)
}
