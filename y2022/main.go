package main

import (
	"context"
	"fmt"
	"os"

	"github.com/espang/aoc/aoc"
	alt_aoc "github.com/espang/aoc/y2022/go/aoc"
	"github.com/espang/aoc/y2022/go/aoc/day1"
	"github.com/espang/aoc/y2022/go/aoc/day10"
	"github.com/espang/aoc/y2022/go/aoc/day11"
	"github.com/espang/aoc/y2022/go/aoc/day12"
	"github.com/espang/aoc/y2022/go/aoc/day13"
	"github.com/espang/aoc/y2022/go/aoc/day14"
	"github.com/espang/aoc/y2022/go/aoc/day15"
	"github.com/espang/aoc/y2022/go/aoc/day16"
	"github.com/espang/aoc/y2022/go/aoc/day17"
	"github.com/espang/aoc/y2022/go/aoc/day18"
	"github.com/espang/aoc/y2022/go/aoc/day19"
	"github.com/espang/aoc/y2022/go/aoc/day20"
	"github.com/espang/aoc/y2022/go/aoc/day21"
	"github.com/espang/aoc/y2022/go/aoc/day22"
	"github.com/espang/aoc/y2022/go/aoc/day23"
	"github.com/espang/aoc/y2022/go/aoc/day24"
	"github.com/espang/aoc/y2022/go/aoc/day25"
	"github.com/espang/aoc/y2022/go/aoc/day9"
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
		fname := aoc.MustMakeInputAvailable(context.TODO(), 2022, 1)
		do(fname, day1.Part1, day1.Part2)
	case "day7":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 7), alt_aoc.Day7Part1, alt_aoc.Day7Part2)
	case "day8":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 8), alt_aoc.Day8Part1, alt_aoc.Day8Part2)
	case "day9":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 9), day9.Part1, day9.Part2)
	case "day10":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 10), day10.Part1, day10.Part2)
	case "day11":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 11), day11.Part1, day11.Part2)
	case "day12":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 12), day12.Part1, day12.Part2)
	case "day13":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 13), day13.Part1, day13.Part2)
	case "day14":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 14), day14.Part1, day14.Part2)
	case "day15":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 15), day15.Part1, day15.Part2)
	case "day16":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 16), day16.Part1, day16.Part2)
	case "day17":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 17), day17.Part1, day17.Part2)
	case "day18":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 18), day18.Part1, day18.Part2)
	case "day19":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 19), day19.Part1, day19.Part2)
	case "day20":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 20), day20.Part1, day20.Part2)
	case "day21":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 21), day21.Part1, day21.Part2)
	case "day22":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 22), day22.Part1, day22.Part2)
	case "day23":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 23), day23.Part1, day23.Part2)
	case "day24":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 24), day24.Part1, day24.Part2)
	case "day25":
		do(aoc.MustMakeInputAvailable(context.TODO(), 2022, 25), day25.Part1, day25.Part2)
	default:
		fmt.Printf("no solution for %q\n", day)
		os.Exit(1)
	}
}
