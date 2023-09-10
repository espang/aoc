package day12

import "strings"

var testinput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

type Point struct {
	x, y int
}

func ParseInput(input string) (start Point, end Point, heights [][]int) {
	lines := strings.Split(input, "\n")
	for rowIdx, line := range lines {
		row := make([]int, 0, len(line))
		for colIdx, c := range line {
			switch c {
			case 'S':
				start.x = rowIdx
				start.y = colIdx
				row = append(row, 1)
			case 'E':
				end.x = rowIdx
				end.y = colIdx
				row = append(row, 26)
			default:
				v := 1 + int(c) - int('a')
				row = append(row, v)
			}
		}
		heights = append(heights, row)
	}
	return start, end, heights
}
