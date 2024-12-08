package main

import (
	"context"
	"fmt"
	"os"

	"github.com/espang/aoc/aoc"
	"github.com/espang/aoc/y2024/go/aoc/day4"
	"github.com/espang/aoc/y2024/go/aoc/day5"
	"github.com/espang/aoc/y2024/go/aoc/day6"
	"github.com/espang/aoc/y2024/go/aoc/day7"
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
	case "day4":
		fname := aoc.MustMakeInputAvailable(context.TODO(), 2024, 4)
		do(fname, day4.Part1, day4.Part2)
	case "day5":
		fname := aoc.MustMakeInputAvailable(context.TODO(), 2024, 5)
		do(fname, day5.Part1, day5.Part2)
	case "day6":
		fname := aoc.MustMakeInputAvailable(context.TODO(), 2024, 6)
		do(fname, day6.Part1, day6.Part2)
	case "day7":
		fname := aoc.MustMakeInputAvailable(context.TODO(), 2024, 7)
		do(fname, day7.Part1, day7.Part2)
	default:
		fmt.Printf("no solution for %q\n", day)
		os.Exit(1)
	}
}
