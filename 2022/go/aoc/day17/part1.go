package day17

import (
	"container/list"
	"fmt"
)

// ##
// ##

// 2 left
// 3 highest

type Point struct{ X, Y int }

func p(x, y int) Point { return Point{X: x, Y: y} }

type PointSet map[Point]struct{}

func FromPoints(ps ...Point) PointSet {
	set := PointSet{}
	for _, p := range ps {
		set.Add(p)
	}
	return set
}
func (ps *PointSet) Add(p Point)          { (*ps)[p] = struct{}{} }
func (ps PointSet) Contains(p Point) bool { _, ok := ps[p]; return ok }
func (ps PointSet) Move(x, y int) PointSet {
	set := PointSet{}
	for p := range ps {
		p.X += x
		p.Y += y
		set.Add(p)
	}
	return set
}
func (ps PointSet) AnyAt(x int) bool {
	for p := range ps {
		if p.X == x {
			return true
		}
	}
	return false
}
func (ps PointSet) Highest() int {
	var max int
	for p := range ps {
		if y := p.Y; y > max {
			max = y
		}
	}
	return max
}

func (ps *PointSet) AddAll(ps2 PointSet) {
	for p := range ps2 {
		ps.Add(p)
	}
}

func Overlap(smallSet, largeSet PointSet) bool {
	for p := range smallSet {
		if largeSet.Contains(p) {
			return true
		}
	}
	return false
}

type Chamber struct {
	Rocks     PointSet
	JetStream JetStream
	Highest   int
	Item      *list.List
}

func NewChamber(js JetStream) Chamber {
	items := list.New()
	// ..####
	items.PushBack(FromPoints(p(2, 0), p(3, 0), p(4, 0), p(5, 0)))
	// ...#.
	// ..###
	// ...#.
	items.PushBack(FromPoints(p(3, 2), p(2, 1), p(3, 1), p(4, 1), p(3, 0)))
	// ....#
	// ....#
	// ..###
	items.PushBack(FromPoints(p(4, 2), p(4, 1), p(2, 0), p(3, 0), p(4, 0)))
	// ..#
	// ..#
	// ..#
	// ..#
	items.PushBack(FromPoints(p(2, 3), p(2, 2), p(2, 1), p(2, 0)))
	// ..##
	// ..##
	items.PushBack(FromPoints(p(2, 1), p(3, 1), p(2, 0), p(3, 0)))
	// the floor
	rocks := PointSet{}
	for x := 0; x <= 6; x++ {
		rocks.Add(p(x, -1))
	}
	return Chamber{
		JetStream: js,
		Item:      items,
		Rocks:     rocks,
	}
}

func (c *Chamber) DropNextRock() {
	// Take the next rock from the list and push it to the back and move it to the right position
	rock := c.Item.PushBack(c.Item.Remove(c.Item.Front())).Value.(PointSet).Move(0, c.Highest+3)
	var didNotMove bool
	for !didNotMove {
		direction := c.JetStream.Next()
		switch direction {
		case left:
			if !rock.AnyAt(0) {
				rock = rock.Move(-1, 0)
				if Overlap(rock, c.Rocks) {
					rock = rock.Move(1, 0)
				}
			}
		case right:
			if !rock.AnyAt(6) {
				rock = rock.Move(1, 0)
				if Overlap(rock, c.Rocks) {
					rock = rock.Move(-1, 0)
				}
			}
		}
		// fall
		rock = rock.Move(0, -1)
		if Overlap(rock, c.Rocks) {
			rock = rock.Move(0, 1)
			didNotMove = true
		}
	}
	c.Rocks.AddAll(rock)
	if rockHighest := rock.Highest() + 1; rockHighest > c.Highest {
		c.Highest = rockHighest
	}
}

func (c Chamber) Show() {
	fmt.Println()
	for y := c.Highest; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			if c.Rocks.Contains(p(x, y)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

func Part1(input string) {
	chamber := NewChamber(Parse(input))
	for step := 0; step < 2_022; step++ {
		chamber.DropNextRock()
		// chamber.Show()
	}
	fmt.Println(chamber.Highest)
}
