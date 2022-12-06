package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sumOfGroup(group []int) int {
	total := 0
	for _, val := range group {
		total += val
	}
	return total
}

func sumOfGroups(groups [][]int) []int {
	ret := make([]int, 0, len(groups))
	for _, group := range groups {
		ret = append(ret, sumOfGroup(group))
	}
	return ret
}

func mustMaxOf(vs []int) int {
	// this panics on purpose
	max := vs[0]
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}

func inputToCaloriesByElf(input string) []int {
	lines := strings.Split(input, "\n")
	groups := [][]int{}
	group := []int{}
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = []int{}
		} else {
			i, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			group = append(group, int(i))
		}
	}
	if lines[len(lines)-1] != "" {
		groups = append(groups, group)
	}
	return sumOfGroups(groups)
}

func Day1Part1(input string) {
	caloriesCarried := inputToCaloriesByElf(input)
	fmt.Print(mustMaxOf(caloriesCarried))
}

func Day1Part2(input string) {
	caloriesCarried := inputToCaloriesByElf(input)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesCarried)))
	val := caloriesCarried[0] + caloriesCarried[1] + caloriesCarried[2]
	fmt.Print(val)
}
