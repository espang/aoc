package day24

import (
	"fmt"

	"github.com/espang/aoc/y2022/go/aoc"
)

type Position struct {
	row, col int
}

type Tornado struct {
	pos       Position
	direction rune
}

type State struct {
	pos   Position
	turns int
}

func moveTornados(width, height int, tornados aoc.Set[Tornado]) aoc.Set[Tornado] {
	newT := aoc.Set[Tornado]{}
	for t := range tornados {
		newPos := t.pos
		switch t.direction {
		case '>':
			newPos.col++
			if newPos.col == width-1 {
				newPos.col = 1
			}

		case '<':
			newPos.col--
			if newPos.col == 0 {
				newPos.col = width - 2
			}

		case '^':
			newPos.row--
			if newPos.row == 0 {
				newPos.row = height - 2
			}

		case 'v':
			newPos.row++
			if newPos.row == height-1 {
				newPos.row = 1
			}

		}
		newT.Add(Tornado{pos: newPos, direction: t.direction})
	}
	return newT
}

func occupiedPositions(tornados aoc.Set[Tornado]) aoc.Set[Position] {
	positions := aoc.Set[Position]{}
	for t := range tornados {
		positions.Add(t.pos)
	}
	return positions
}

func potentialPositions(pos Position, start, end Position, width, height int) []Position {
	deltas := []Position{{-1, 0}, {1, 0}, {0, 0}, {0, -1}, {0, 1}}
	positions := []Position{}
	for _, delta := range deltas {
		delta.row += pos.row
		delta.col += pos.col
		if delta == start || delta == end {
			positions = append(positions, delta)
			continue
		}
		if delta.row <= 0 || delta.row >= height-1 {
			continue
		}
		if delta.col == 0 || delta.col == width-1 {
			continue
		}
		positions = append(positions, delta)
	}
	return positions
}

func minimumEscape(start, end Position, width, height int, tornados aoc.Set[Tornado]) int {
	if start == end {
		return 0
	}
	currentPositions := aoc.Set[Position]{}
	currentPositions.Add(start)
	currentTornados := tornados
	for turn := 0; ; turn++ {
		if len(currentPositions) == 0 {
			return -1
		}
		currentTornados = moveTornados(width, height, currentTornados)
		occupied := occupiedPositions(currentTornados)
		nextPositions := aoc.Set[Position]{}
		for pos := range currentPositions {
			moves := potentialPositions(pos, start, end, width, height)
			for _, move := range moves {
				if !occupied.IsMember(move) {
					if move == end {
						return turn + 1
					}
					nextPositions.Add(move)
				}
			}
		}
		currentPositions = nextPositions
	}
	return -1
}

func Part1(input string) {
	board := Parse(input)
	start := Position{row: 0}
	end := Position{row: len(board) - 1}
	for col, c := range board[0] {
		if c == '.' {
			start.col = col
			break
		}
	}
	for col, c := range board[len(board)-1] {
		if c == '.' {
			end.col = col
			break
		}
	}
	tornados := aoc.Set[Tornado]{}
	for rowIdx, row := range board {
		for colIdx, cell := range row {
			switch cell {
			case '>', '<', '^', 'v':
				tornados.Add(Tornado{
					pos:       Position{row: rowIdx, col: colIdx},
					direction: cell,
				})
			}
		}
	}
	width := len(board[0])
	height := len(board)
	fmt.Println(minimumEscape(start, end, width, height, tornados))
}
