package day18

import (
	"fmt"
)

type CoordSet map[Coord]struct{}

func From(coords ...Coord) CoordSet {
	cs := make(CoordSet, len(coords))
	for _, c := range coords {
		cs.Add(c)
	}
	return cs
}

func (cs CoordSet) Add(c Coord) CoordSet {
	cs[c] = struct{}{}
	return cs
}

func (cs CoordSet) Contains(c Coord) bool {
	_, ok := cs[c]
	return ok
}

func (cs CoordSet) AddAll(cs2 CoordSet) CoordSet {
	for k := range cs2 {
		cs.Add(k)
	}
	return cs
}

func neighboursOf(c Coord) []Coord {
	deltas := []Coord{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	neighbours := make([]Coord, 0, 6)
	for _, delta := range deltas {
		neighbours = append(neighbours,
			Coord{
				X: c.X + delta.X,
				Y: c.Y + delta.Y,
				Z: c.Z + delta.Z,
			})
	}
	return neighbours
}

func countConnections(s CoordSet) map[Coord]int {
	connections := make(map[Coord]int, len(s))
	for c := range s {
		neighbours := neighboursOf(c)
		connectionsOfC := 0
		for _, n := range neighbours {
			if s.Contains(n) {
				connectionsOfC++
			}
		}
		connections[c] = connectionsOfC
	}
	return connections
}

func countSides(cs CoordSet) int {
	totalSides := 0
	for _, nconns := range countConnections(cs) {
		totalSides += 6 - nconns
	}
	return totalSides
}

func Part1(input string) {
	coordinates := Parse(input)
	totalSides := countSides(From(coordinates...))
	fmt.Println(totalSides)
}
