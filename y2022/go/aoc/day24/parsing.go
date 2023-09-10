package day24

import "strings"

var testinput = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func Parse(input string) [][]rune {
	matrix := [][]rune{}
	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, []rune(row))
	}
	return matrix
}
