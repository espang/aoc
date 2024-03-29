package main

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/espang/aoc/y2017/day21"
	"github.com/espang/aoc/y2017/day23"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		slog.Warn("please provide the day you want to solve as a command line argument")
		os.Exit(1)
	}
	// ignore additional arguments
	day := strings.ToLower(args[1])
	var part1 func() (string, error)
	var part2 func() (string, error)
	switch day {
	case "day21":
		part1 = day21.Part1
		part2 = day21.Part2
	case "day23":
		part1 = day23.Part1
		part2 = day23.Part2
	default:
		slog.Warn("couldn't find solution for the given day, expect this format `dayN` for day N and N can be any number between 1 and 25", slog.String("day-value", day))
		os.Exit(1)
	}

	start := time.Now()
	result, err := part1()
	if err != nil {
		slog.Warn("failed to handle part1", slog.String("day", day), slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("result for part 1 is", slog.String("result", result), slog.Duration("took", time.Since(start)))
	start = time.Now()
	result, err = part2()
	if err != nil {
		slog.Warn("failed to handle part2", slog.String("day", day), slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("result for part 2 is", slog.String("result", result), slog.Duration("took", time.Since(start)))
}
