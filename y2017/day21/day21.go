package day21

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/espang/aoc/aoc"
)

type BluePrints struct {
	twoToThree            [][]int
	twoTransformIndexes   map[int]int
	threeToFour           [][]int
	threeTransformIndexes map[int]int
}

func (b *BluePrints) Add(from []int, to []int) {
	if len(from) == 4 {
		setBefore := aoc.NewSet(aoc.MapKeys(b.twoTransformIndexes)...)
		index := len(b.twoToThree)
		b.twoToThree = append(b.twoToThree, to)
		for _, m := range allPermutations(from) {
			val := hash(m)
			if setBefore.Contains(val) {
				fmt.Println(from, " --> ", to)
				panic("add - seen value before")
			}
			// point to the 3x3 to transform it too
			b.twoTransformIndexes[val] = index
		}
		return
	}
	if len(from) == 9 {
		setBefore := aoc.NewSet(aoc.MapKeys(b.threeTransformIndexes)...)
		index := len(b.threeToFour)
		b.threeToFour = append(b.threeToFour, to)

		for _, m := range allPermutations(from) {
			val := hash(m)
			if setBefore.Contains(val) {
				fmt.Println(from, " --> ", to)
				panic("add - seen value before")
			}
			// point to the 3x3 to transform it too
			b.threeTransformIndexes[val] = index
		}
		return
	}
	panic("blueprints::Add")
}

func (b *BluePrints) Transform(mat []int) []int {
	h := hash(mat)
	if len(mat) == 4 {
		index, ok := b.twoTransformIndexes[h]
		if !ok {
			panic("Transform 2x2")
		}
		return b.twoToThree[index]
	}
	if len(mat) == 9 {
		index, ok := b.threeTransformIndexes[h]
		if !ok {
			panic("Transform 3x3")
		}
		return b.threeToFour[index]
	}
	panic("invalid length: Transform")
}

// hash transforms the square into an integer that
// can be added to sets or maps to look it up
func hash(mat []int) int {
	var result int
	for i, v := range mat {
		if v == 1 {
			result = aoc.SetBit(result, i)
		}
	}
	return result
}

func allPermutations(mat []int) [][]int {
	var permutations [][]int

	toChange := slices.Clone(mat)
	for i := 0; i < 4; i++ {
		toChange = rotate(toChange)
		permutations = append(permutations, slices.Clone(toChange))
	}

	toChange = flip(slices.Clone(mat))
	for i := 0; i < 4; i++ {
		toChange = rotate(toChange)
		permutations = append(permutations, slices.Clone(toChange))
	}

	// toChange = flip(slices.Clone(mat))
	// toChange = rotate(toChange)
	// toChange = flip(toChange)
	// for i := 0; i < 4; i++ {
	// 	permutations = append(permutations, slices.Clone(toChange))
	// 	toChange = rotate(toChange)

	// }

	return permutations
}

func rotate(matrix []int) []int {
	if len(matrix) == 4 {
		// 0 1  --> 2 0
		// 2 3  --> 3 1
		return []int{
			matrix[2], matrix[0],
			matrix[3], matrix[1],
		}
	}
	if len(matrix) == 9 {
		// 0 1 2 --> 6 3 0
		// 3 4 5 --> 7 4 1
		// 6 7 8 --> 8 5 2
		return []int{
			matrix[6], matrix[3], matrix[0],
			matrix[7], matrix[4], matrix[1],
			matrix[8], matrix[5], matrix[2],
		}
	}
	panic("rotate")
}

func flip(matrix []int) []int {
	if len(matrix) == 4 {
		// 0 1  --> 2 3
		// 2 3  --> 0 1
		return []int{
			matrix[2], matrix[3],
			matrix[0], matrix[1],
		}
	}
	if len(matrix) == 9 {
		// 0 1 2 --> 6 7 8
		// 3 4 5 --> 3 4 5
		// 6 7 8 --> 0 1 2
		return []int{
			matrix[6], matrix[7], matrix[8],
			matrix[3], matrix[4], matrix[5],
			matrix[0], matrix[1], matrix[2],
		}
	}
	panic("flip")
}

func parseLine(line string) ([]int, []int) {
	splitted := strings.Split(line, " => ")
	return parse(splitted[0]),
		parse(splitted[1])
}

func parse(s string) []int {
	var mat []int
	for _, c := range s {
		switch c {
		case '.':
			mat = append(mat, 0)
		case '#':
			mat = append(mat, 1)
		}
	}
	return mat
}

type State struct {
	m    *aoc.Set[aoc.Coordinate]
	size int
}

func (s State) Show() {
	var loc aoc.Coordinate
	for y := 0; y < s.size; y++ {
		for x := 0; x < s.size; x++ {
			loc.X = x
			loc.Y = y
			if s.m.Contains(loc) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (s State) next2(b *BluePrints) State {
	newMat := aoc.NewSet[aoc.Coordinate]()
	var loc aoc.Coordinate
	for x := 0; x < s.size; x += 2 {
		for y := 0; y < s.size; y += 2 {
			mat := make([]int, 4)
			loc.X = x
			loc.Y = y
			for i, delta := range []aoc.Coordinate{
				{0, 0}, {1, 0},
				{0, 1}, {1, 1},
			} {
				if s.m.Contains(loc.Plus(delta)) {
					mat[i] = 1
				}
			}
			loc.X = (x / 2) * 3
			loc.Y = (y / 2) * 3
			out := b.Transform(mat)
			for i, v := range out {
				delta := aoc.Coordinate{
					X: i % 3,
					Y: i / 3,
				}
				if v == 1 {
					newMat.Add(loc.Plus(delta))
				}
			}
		}
	}
	return State{
		size: (s.size / 2) * 3,
		m:    newMat,
	}
}

func (s State) next3(b *BluePrints) State {
	newMat := aoc.NewSet[aoc.Coordinate]()
	var loc aoc.Coordinate
	for x := 0; x < s.size; x += 3 {
		for y := 0; y < s.size; y += 3 {
			mat := make([]int, 9)
			loc.X = x
			loc.Y = y
			for i, delta := range []aoc.Coordinate{
				{0, 0}, {1, 0}, {2, 0},
				{0, 1}, {1, 1}, {2, 1},
				{0, 2}, {1, 2}, {2, 2},
			} {
				if s.m.Contains(loc.Plus(delta)) {
					mat[i] = 1
				}
			}
			loc.X = (x / 3) * 4
			loc.Y = (y / 3) * 4
			out := b.Transform(mat)
			for i, v := range out {
				delta := aoc.Coordinate{
					X: i % 4,
					Y: i / 4,
				}
				if v == 1 {
					newMat.Add(loc.Plus(delta))
				}
			}
		}
	}
	return State{
		size: (s.size / 3) * 4,
		m:    newMat,
	}
}

func (s State) Next(b *BluePrints) State {
	if s.size%2 == 0 {
		return s.next2(b)
	} else if s.size%3 == 0 {
		return s.next3(b)
	}
	panic("Next")
}

func part1(content string, n int) (string, error) {
	lines := aoc.SplitBy(content, aoc.Newline)

	bluePrints := &BluePrints{
		twoTransformIndexes:   map[int]int{},
		threeTransformIndexes: map[int]int{},
	}
	f := func(line string) {
		bluePrints.Add(parseLine(line))
	}
	// Reads the input and adds the transformation of each line
	// to the blueprints.
	aoc.Run(f, lines)

	initialState := State{
		size: 3,
		// .#.
		// ..#
		// ###
		m: aoc.NewSet(
			aoc.Coordinate{X: 1, Y: 0},
			aoc.Coordinate{X: 2, Y: 1},
			aoc.Coordinate{X: 0, Y: 2},
			aoc.Coordinate{X: 1, Y: 2},
			aoc.Coordinate{X: 2, Y: 2},
		),
	}
	state := initialState
	for i := 0; i < n; i++ {
		// state.Show()
		state = state.Next(bluePrints)
	}
	// state.Show()
	return strconv.Itoa(state.m.Len()), nil
}

func Part1() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2017, 21)
	content := aoc.LoadFrom(filename)
	return part1(content, 5)
}

func Part2() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2017, 21)
	content := aoc.LoadFrom(filename)
	return part1(content, 18)
}
