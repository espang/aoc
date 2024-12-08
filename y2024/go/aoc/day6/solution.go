package day6

import (
	"fmt"

	"github.com/espang/aoc/aoc"
)

// Parse parses the input and returns the Board and Player.
// panics when the data is unexpected.
func Parse(s string) (Board, Player) {
	rows := aoc.SplitByLine(s)
	ncols := len(rows[0])
	if !aoc.All(func(row string) bool { return len(row) == ncols }, rows) {
		panic("unexpected input shape")
	}
	var obstacles [][]bool
	var player *Player
	for y, row := range rows {
		var obstaclesInRow []bool
		for x, char := range row {
			switch char {
			case '#':
				obstaclesInRow = append(obstaclesInRow, true)
			case '.':
				obstaclesInRow = append(obstaclesInRow, false)
			default:
				dir := aoc.ParseDirection(char)
				if dir == 0 {
					panic("unexpected input char")
				}
				obstaclesInRow = append(obstaclesInRow, false)

				if player != nil {
					panic("unexpected player")
				}
				player = &Player{
					dir: dir,
					at:  aoc.Coordinate{X: x, Y: -y},
				}
			}
		}
		obstacles = append(obstacles, obstaclesInRow)
	}
	return Board{
		obstacles: obstacles,
		ncols:     ncols,
		nrows:     len(obstacles),
	}, *player
}

type Player struct {
	at  aoc.Coordinate
	dir aoc.Direction
}

func (p Player) DoStep(b Board) Player {
	next := aoc.Move(p.dir)(p.at)
	if b.IsObstacle(next) {
		p.dir = p.dir.Right()
	} else {
		p.at = next
	}
	return p
}

type Board struct {
	obstacles [][]bool
	ncols     int
	nrows     int
}

func (b Board) IsOut(c aoc.Coordinate) bool {
	return (c.Y > 0 || -c.Y >= b.nrows) ||
		(c.X < 0 || c.X >= b.ncols)
}

func (b Board) IsObstacle(c aoc.Coordinate) bool {
	if b.IsOut(c) {
		return false
	}
	return b.obstacles[-c.Y][c.X]
}

func (b *Board) SetObstacle(c aoc.Coordinate) {
	b.obstacles[-c.Y][c.X] = true
}

func (b *Board) RemoveObstacle(c aoc.Coordinate) {
	b.obstacles[-c.Y][c.X] = false
}

func Part1(s string) {
	board, player := Parse(s)

	visited := aoc.EmptySet[aoc.Coordinate]()
	for !board.IsOut(player.at) {
		visited.Add(player.at)
		player = player.DoStep(board)
	}
	fmt.Println(visited.Len())
}

func DoesLoop(board Board, player Player) bool {
	visited := aoc.EmptySet[Player]()
	for !board.IsOut(player.at) {
		if visited.Contains(player) {
			return true
		}
		visited.Add(player)
		player = player.DoStep(board)
	}
	return false
}

func Part2(s string) {
	board, player := Parse(s)
	total := 0
	for x := 0; x < board.ncols; x++ {
		for y := 0; y > -board.nrows; y-- {
			c := aoc.Coordinate{X: x, Y: y}
			if !board.IsObstacle(c) {
				board.SetObstacle(c)
				if DoesLoop(board, player) {
					total++
				}
				board.RemoveObstacle(c)
			}
		}
	}
	fmt.Println(total)
}
