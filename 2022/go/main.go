package main

import (
	"aoc/aoc"
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
	default:
		fmt.Printf("no solution for %q\n", day)
		os.Exit(1)
	}
}
