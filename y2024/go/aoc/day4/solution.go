package day4

import (
	"fmt"

	"github.com/espang/aoc/aoc"
)

var (
	// rows; cols
	right = [3][2]int{{0, 1}, {0, 2}, {0, 3}}
	left  = [3][2]int{{0, -1}, {0, -2}, {0, -3}}
	up    = [3][2]int{{-1, 0}, {-2, 0}, {-3, 0}}
	down  = [3][2]int{{1, 0}, {2, 0}, {3, 0}}
	// diagonals
	rightDown = [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	leftDown  = [3][2]int{{1, -1}, {2, -2}, {3, -3}}
	rightUp   = [3][2]int{{-1, 1}, {-2, 2}, {-3, 3}}
	leftUp    = [3][2]int{{-1, -1}, {-2, -2}, {-3, -3}}
)

type Board struct {
	values [][]rune
	counts [][]int
}

func (b Board) HasValueAt(row, col int) bool {
	return row >= 0 && row < len(b.values) &&
		len(b.values) > 0 &&
		col >= 0 && col < len(b.values[0])
}

func (b Board) ValueAt(row, col int) (rune, bool) {
	if b.HasValueAt(row, col) {
		return b.values[row][col], true
	}
	return 0, false
}

func (b Board) ValueAtIs(row, col int, r rune) bool {
	val, ok := b.ValueAt(row, col)
	return ok && val == r
}

// check takes a current position and returns a function that takes
// 3 moves and checks wether the charcter after move 1 is M the character
// after move 2 is A and the character after move 3 is S.
func (b Board) check(row, col int) func([3][2]int) int {
	return func(directions [3][2]int) int {
		if b.ValueAtIs(row+directions[0][0], col+directions[0][1], 'M') &&
			b.ValueAtIs(row+directions[1][0], col+directions[1][1], 'A') &&
			b.ValueAtIs(row+directions[2][0], col+directions[2][1], 'S') {
			return 1
		}
		return 0
	}
}

func (b Board) CalculateCounts(count func(int, int) int) int {
	var total int
	for rowI, row := range b.values {
		for colI := range row {
			b.counts[rowI][colI] = count(rowI, colI)
			total += b.counts[rowI][colI]
		}
	}
	return total
}

func parse(input string) Board {
	board := Board{}
	lines := aoc.SplitByLine(input)
	for _, line := range lines {
		var rowValues []rune
		for _, r := range line {
			rowValues = append(rowValues, r)
		}
		board.values = append(board.values, rowValues)
		board.counts = append(board.counts, make([]int, len(rowValues)))
	}
	return board
}

func Part1(input string) {
	b := parse(input)
	calcCellValuePart := func(row, col int) int {
		if !b.ValueAtIs(row, col, 'X') {
			return 0
		}
		checkF := b.check(row, col)
		b.counts[row][col] = checkF(right) + checkF(left) + checkF(up) + checkF(down) +
			checkF(rightUp) + checkF(rightDown) + checkF(leftUp) + checkF(leftDown)
		return b.counts[row][col]
	}

	total := b.CalculateCounts(calcCellValuePart)
	fmt.Printf("%d", total)
}

func Part2(input string) {
	b := parse(input)
	calcCellValuePart := func(row, col int) int {
		if !b.ValueAtIs(row, col, 'A') {
			return 0
		}

		leftTopToRightBottom := (b.ValueAtIs(row-1, col-1, 'M') && b.ValueAtIs(row+1, col+1, 'S')) ||
			(b.ValueAtIs(row-1, col-1, 'S') && b.ValueAtIs(row+1, col+1, 'M'))
		leftBottomToRightTop := (b.ValueAtIs(row+1, col-1, 'M') && b.ValueAtIs(row-1, col+1, 'S')) ||
			(b.ValueAtIs(row+1, col-1, 'S') && b.ValueAtIs(row-1, col+1, 'M'))
		if leftTopToRightBottom && leftBottomToRightTop {
			return 1
		}
		return 0
	}
	total := b.CalculateCounts(calcCellValuePart)
	fmt.Printf("%d", total)
}
