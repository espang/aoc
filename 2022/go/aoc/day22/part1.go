package day22

import "fmt"

func ExecuteInstructions(instructions []any, grid Grid) (int, int, Orientation) {
	orientation := Right
	row := 0
	col := -1
	for c := 0; c < len(grid[0]); c++ {
		if grid.At(row, c) == Open {
			col = c
			break
		}
	}
	if col == -1 {
		panic("first open is a wall - this is likely allowed here and should be handled if necessary")
	}
	for _, instruction := range instructions {
		switch val := instruction.(type) {
		case Move:
			row, col = grid.Handle(val, row, col, orientation)
		case Rotate:
			orientation = orientation.Rotate(val)
		}
	}
	return row, col, orientation
}

func Part1(input string) {
	grid, instructions := Parse(input)
	row, col, orientation := ExecuteInstructions(instructions, grid)
	result := 1000*(row+1) + 4*(col+1) + int(orientation)
	fmt.Println(result)
}
