package aoc

type MoveFn func(Coordinate) Coordinate

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Plus(c2 Coordinate) Coordinate {
	return Coordinate{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

func (c Coordinate) Equal(c2 Coordinate) bool {
	return c.X == c2.X && c.Y == c2.Y
}

func MoveUp(c Coordinate) Coordinate {
	c.Y += 1
	return c
}

func MoveDown(c Coordinate) Coordinate {
	c.Y -= 1
	return c
}

func MoveLeft(c Coordinate) Coordinate {
	c.X -= 1
	return c
}

func MoveRight(c Coordinate) Coordinate {
	c.X += 1
	return c
}

// Move returns a function that moves the given coordinate 1 step
// into Direaction 'd'.
func Move(d Direction) func(Coordinate) Coordinate {
	switch d {
	case Up:
		return MoveUp
	case Down:
		return MoveDown
	case Left:
		return MoveLeft
	case Right:
		return MoveRight
	}
	// ignore invalid input and return invalid direction
	return func(c Coordinate) Coordinate { return c }
}

func ManhattanDistance(c1, c2 Coordinate) int {
	return abs(c1.X-c2.X) + abs(c1.Y-c2.Y)
}

type Direction int

const (
	Up Direction = iota + 1
	Down
	Left
	Right
)

func ParseDirection(r rune) Direction {
	switch r {
	case '>':
		return Right
	case '<':
		return Left
	case '^':
		return Up
	case 'v':
		return Down
	}

	// ignore unexpect input and return invalid direction
	return 0
}

func (d Direction) Left() Direction {
	switch d {
	case Up:
		return Left
	case Left:
		return Down
	case Down:
		return Right
	case Right:
		return Up
	}
	// ignore invalid input and return invalid direction
	return d
}

func (d Direction) Right() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	// ignore invalid input and return invalid direction
	return d
}
