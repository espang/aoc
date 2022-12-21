package main

import (
	"aoc/aoc"
	"aoc/aoc/day10"
	"aoc/aoc/day11"
	"aoc/aoc/day12"
	"aoc/aoc/day13"
	"aoc/aoc/day14"
	"aoc/aoc/day15"
	"aoc/aoc/day16"
	"aoc/aoc/day17"
	"aoc/aoc/day18"
	"aoc/aoc/day19"
	"aoc/aoc/day20"
	"aoc/aoc/day21"
	"aoc/aoc/day9"
	"fmt"
	"os"
)

func readInput(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func do(filename string, part1 func(string), part2 func(string)) {
	input, err := readInput(filename)
	if err != nil {
		fmt.Printf("error reading input file: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("Part1: ")
	part1(input)
	fmt.Println("")
	fmt.Print("Part2: ")
	part2(input)
	fmt.Println("")
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Expects exactly one argument to select the day to solve. for example 'day1'.")
		os.Exit(1)
	}

	day := os.Args[1]
	switch day {
	case "day1":
		do("../inputs/input1.txt", aoc.Day1Part1, aoc.Day1Part2)
	case "day7":
		do("../inputs/input7.txt", aoc.Day7Part1, aoc.Day7Part2)
	case "day8":
		do("../inputs/input8.txt", aoc.Day8Part1, aoc.Day8Part2)
	case "day9":
		do("../inputs/input9.txt", day9.Part1, day9.Part2)
	case "day10":
		do("../inputs/input10.txt", day10.Part1, day10.Part2)
	case "day11":
		do("../inputs/input11.txt", day11.Part1, day11.Part2)
	case "day12":
		do("../inputs/input12.txt", day12.Part1, day12.Part2)
	case "day13":
		do("../inputs/input13.txt", day13.Part1, day13.Part2)
	case "day14":
		do("../inputs/input14.txt", day14.Part1, day14.Part2)
	case "day15":
		do("../inputs/input15.txt", day15.Part1, day15.Part2)
	case "day16":
		do("../inputs/input16.txt", day16.Part1, day16.Part2)
	case "day17":
		do("../inputs/input17.txt", day17.Part1, day17.Part2)
	case "day18":
		do("../inputs/input18.txt", day18.Part1, day18.Part2)
	case "day19":
		do("../inputs/input19.txt", day19.Part1, day19.Part2)
	case "day20":
		do("../inputs/input20.txt", day20.Part1, day20.Part2)
	case "day21":
		do("../inputs/input21.txt", day21.Part1, day21.Part2)
	default:
		fmt.Printf("no solution for %q\n", day)
		os.Exit(1)
	}
}
