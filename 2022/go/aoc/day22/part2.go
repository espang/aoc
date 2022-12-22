package day22

import (
	"fmt"
)

type State struct {
	row         int
	col         int
	orientation Orientation
}

func s(row, col int, o Orientation) State {
	return State{
		row:         row,
		col:         col,
		orientation: o,
	}
}

func foldsOfTestInput() map[State]State {
	// testinput
	// .   .    Front  .
	// Top Left Bottom .
	// .   .    Back   Right
	folds := map[State]State{}

	// Front to Top
	for col := 8; col < 12; col++ {
		toCol := 11 - col
		folds[s(-1, col, Up)] = s(4, toCol, Down)
		folds[s(3, toCol, Up)] = s(0, col, Down)
	}
	for row := 0; row < 4; row++ {
		// ** Front to Left
		// 0,8, LEFT -> 4,4 -> Down
		// 1,8, LEFT -> 4,5 -> Down
		folds[s(row, 7, Left)] = s(4, 4+row, Down)
		// 4,8, UP -> 0,4, Right
		folds[s(3, row+4, Up)] = s(row, 8, Right)

		// ** Front to Right ((8,12) -> (11, 15))
		// 0,11, Right -> 11, 15, Left
		// 1,11, Right -> 10, 15, Left
		folds[s(row, 12, Right)] = s(11, 15, Left)
		// 8, 15, Right -> 3, 11, Left
		// 11, 15, Right -> 0, 11, Left
		folds[s(8+row, 16, Right)] = s(3-row, 11, Left)

		// ** Bottom to right
		// 4, 11, Right -> 8, 15, Down
		// 5, 11, Right -> 8, 14, Down
		folds[s(4+row, 12, Right)] = s(8, 15-row, Down)
		// 8,15, Up -> 4, 11, Left
		// 8,14, Up -> 5, 11, Left
		folds[s(7, 15-row, Up)] = s(4+row, 11, Left)

		// ** Left to Back
		// 7, 4, Down -> 11, 8 Right
		// 7, 5, Down -> 10, 8 Right
		folds[s(8, 4+row, Down)] = s(11-row, 8, Right)
		// 11, 8, Left -> 7, 4 Up
		folds[s(11-row, 7, Left)] = s(7, 4+row, Up)

		// ** Top to Back
		// 7, 0, Down -> 11, 11, Up
		// 7, 1, Down -> 11, 10, Up
		folds[s(8, row, Down)] = s(11, 11-row, Up)
		// 11, 11, Down -> 7, 0 Up
		// 11, 10, Down -> 7, 1 Up
		folds[s(12, 11-row, Down)] = s(7, row, Up)

		// ** Top To Right
		// 4, 0, Left -> 11, 15 Up
		// 7, 0, Left -> 11, 12, Up
		folds[s(4+row, -1, Left)] = s(11, 15-row, Up)
		folds[s(12, 15-row, Down)] = s(4+row, 0, Right)
	}

	for row := 4; row < 8; row++ {

	}

	return folds
}

func foldOfInput() map[State]State {
	// input
	// .     Front Right
	// .     Bottom
	// Left  Back
	// Top   .
	folds := map[State]State{}
	for i := 0; i < 50; i++ {
		// Front to Top
		folds[s(-1, 50+i, Up)] = s(150+i, 0, Right)
		folds[s(150+i, -1, Left)] = s(0, 50+i, Down)
		// Front to Left
		folds[s(i, 49, Left)] = s(149-i, 0, Right)
		folds[s(149-i, -1, Left)] = s(i, 50, Right)
		// Right to Top
		folds[s(-1, 100+i, Up)] = s(199, i, Up)
		folds[s(200, i, Down)] = s(0, 100+i, Down)
		// Right to Back
		folds[s(i, 150, Right)] = s(149-i, 99, Left)
		folds[s(149-i, 100, Right)] = s(i, 149, Left)
		// Right to Bottom
		folds[s(50, 100+i, Down)] = s(50+i, 99, Left)
		folds[s(50+i, 100, Right)] = s(49, 100+i, Up)
		// Bottom to Left
		folds[s(50+i, 49, Left)] = s(100, i, Down)
		folds[s(99, i, Up)] = s(50+i, 50, Right)
		// Top to Back
		folds[s(150+i, 50, Right)] = s(149, 50+i, Up)
		folds[s(150, 50+i, Down)] = s(150+i, 49, Left)
	}
	return folds
}

func ExecuteInstructions2(instructions []any, grid Grid, folds map[State]State) (int, int, Orientation) {
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
			for i := 0; i < int(val); i++ {
				nextRow, nextCol, nextOrientation := grid.HandleCube(row, col, orientation, folds)
				if nextRow == row && nextCol == col {
					// no move, wall
					break
				}
				row = nextRow
				col = nextCol
				orientation = nextOrientation
			}
		case Rotate:
			orientation = orientation.Rotate(val)
		}
	}
	return row, col, orientation
}

func Part2(input string) {
	// input = testinput
	// folds := foldsOfTestInput()
	folds := foldOfInput()
	grid, instructions := Parse(input)
	row, col, orientation := ExecuteInstructions2(instructions, grid, folds)
	result := 1000*(row+1) + 4*(col+1) + int(orientation)
	fmt.Println(result)
}
