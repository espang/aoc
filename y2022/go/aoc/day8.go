package aoc

import (
	"fmt"
	"strings"
)

func parse8row(s string) []int {
	vs := []int{}
	for _, digit := range s {
		vs = append(vs, int(digit)-int('0'))
	}
	return vs
}

func visibleUp(mat [][]int, row, col int) bool {
	height := mat[row][col]
	for rowi := row - 1; rowi >= 0; rowi-- {
		if mat[rowi][col] >= height {
			return false
		}
	}
	return true
}

func visibleUpN(mat [][]int, row, col int) int {
	height := mat[row][col]
	count := 0
	for rowi := row - 1; rowi >= 0; rowi-- {
		count++
		if mat[rowi][col] >= height {
			return count
		}
	}
	return count
}

func visibleDown(mat [][]int, row, col int) bool {
	height := mat[row][col]
	for rowi := row + 1; rowi < len(mat); rowi++ {
		if mat[rowi][col] >= height {
			return false
		}
	}
	return true
}

func visibleDownN(mat [][]int, row, col int) int {
	height := mat[row][col]
	count := 0
	for rowi := row + 1; rowi < len(mat); rowi++ {
		count++
		if mat[rowi][col] >= height {
			return count
		}
	}
	return count
}

func visibleLeft(mat [][]int, row, col int) bool {
	height := mat[row][col]
	for coli := col - 1; coli >= 0; coli-- {
		if mat[row][coli] >= height {
			return false
		}
	}
	return true
}

func visibleLeftN(mat [][]int, row, col int) int {
	height := mat[row][col]
	count := 0
	for coli := col - 1; coli >= 0; coli-- {
		count++
		if mat[row][coli] >= height {
			return count
		}
	}
	return count
}

func visibleRight(mat [][]int, row, col int) bool {
	height := mat[row][col]
	for coli := col + 1; coli < len(mat[0]); coli++ {
		if mat[row][coli] >= height {
			return false
		}
	}
	return true
}

func visibleRightN(mat [][]int, row, col int) int {
	height := mat[row][col]
	count := 0
	for coli := col + 1; coli < len(mat[0]); coli++ {
		count++
		if mat[row][coli] >= height {
			return count
		}
	}
	return count
}

func isVisible(mat [][]int, row, col int) bool {
	return visibleUp(mat, row, col) || visibleDown(mat, row, col) || visibleLeft(mat, row, col) || visibleRight(mat, row, col)
}

func scenicScore(mat [][]int, row, col int) int {
	return visibleUpN(mat, row, col) *
		visibleDownN(mat, row, col) *
		visibleLeftN(mat, row, col) *
		visibleRightN(mat, row, col)
}

func Day8Part1(input string) {
	lines := strings.Split(input, "\n")
	rows := [][]int{}
	for _, line := range lines {
		rows = append(rows, parse8row(line))
	}
	nrows := len(rows)
	ncols := len(rows[0])
	total := 0
	for row := 0; row < nrows; row++ {
		for col := 0; col < ncols; col++ {
			if isVisible(rows, row, col) {
				total++
			}
		}
	}
	fmt.Println(total)
}
func Day8Part2(input string) {
	lines := strings.Split(input, "\n")
	rows := [][]int{}
	for _, line := range lines {
		rows = append(rows, parse8row(line))
	}
	nrows := len(rows)
	ncols := len(rows[0])
	max := 0
	for row := 0; row < nrows; row++ {
		for col := 0; col < ncols; col++ {
			score := scenicScore(rows, row, col)
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}
