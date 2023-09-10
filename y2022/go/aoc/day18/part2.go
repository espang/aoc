package day18

import (
	"container/list"
	"fmt"
	"math"
)

type Air struct {
	min Coord
	max Coord
}

func NewAir(cs CoordSet) Air {
	minX := math.MaxInt
	minY := math.MaxInt
	minZ := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	maxZ := math.MinInt
	for c := range cs {
		if c.X < minX {
			minX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Z < minZ {
			minZ = c.Z
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.Z > maxZ {
			maxZ = c.Z
		}
	}
	return Air{
		min: Coord{X: minX, Y: minY, Z: minZ},
		max: Coord{X: maxX, Y: maxY, Z: maxZ},
	}
}

func (a Air) isAtAir(c Coord) bool {
	if c.X >= a.max.X || c.Y >= a.max.Y || c.Z >= a.max.Z {
		return true
	}
	if c.X <= a.min.X || c.Y <= a.min.Y || c.Z <= a.min.Z {
		return true
	}
	return false
}

func WayToAir(c Coord, air Air, all CoordSet) bool {
	seen := CoordSet{}.Add(c).AddAll(all)
	queue := list.New()
	queue.PushBack(c)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Coord)
		if air.isAtAir(current) {
			return true
		}
		for _, n := range neighboursOf(current) {
			if seen.Contains(n) {
				continue
			}
			seen.Add(n)
			queue.PushBack(n)
		}
	}
	return false
}

func addAirdropsToSet(set CoordSet) CoordSet {
	air := NewAir(set)
	c := Coord{}
	for x := air.min.X; x <= air.max.X; x++ {
		c.X = x
		for y := air.min.Y; y <= air.max.Y; y++ {
			c.Y = y
			for z := air.min.Z; z <= air.max.Z; z++ {
				c.Z = z
				if set.Contains(c) {
					continue
				}
				if WayToAir(c, air, set) {
					continue
				}
				// treat the airdrop as lava
				set = set.Add(c)
			}
		}
	}
	return set
}

func Part2(input string) {
	totalSides := countSides(addAirdropsToSet(From(Parse(input)...)))
	fmt.Println(totalSides)
}
