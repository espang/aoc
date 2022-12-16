package day16

import (
	"fmt"
)

func solvePart1(start string, maxTime int, valves []Valve) int {
	nodes, rates := nonNullValves(valves)
	distances := distances(nodes, valves)

	var solve func(int, int, NodeSet) int
	solve = func(at int, timeLeft int, seen NodeSet) int {
		var max int
		for to := range nodes {
			if seen.Contains(to) {
				continue
			}
			tl := timeLeft - distances[at][to]
			if tl > 0 {
				v := tl*rates[to] + solve(to, tl, seen.Add(to))
				if max < v {
					max = v
				}
			}

		}
		return max
	}

	var max int
	var seen NodeSet
	for i, node := range nodes {
		travelTime := 1 + ShortestPath(start, node, valves)
		timeLeft := maxTime - travelTime
		pressure := timeLeft * rates[i]
		val := pressure + solve(i, timeLeft, seen.Add(i))
		if val > max {
			max = val
		}
	}
	return max
}

func distances(nodes []string, valves []Valve) [][]int {
	mat := make([][]int, len(nodes))
	for row := 0; row < len(nodes); row++ {
		mat[row] = make([]int, len(nodes))
	}
	for row := 0; row < len(nodes); row++ {
		for col := row + 1; col < len(nodes); col++ {
			start := nodes[row]
			target := nodes[col]
			// 1 is added for the opening of the valve; we will not visit
			// any node twice to open a valve. Any further visit might
			mat[row][col] = 1 + ShortestPath(start, target, valves)
			// mirror the values for easier use
			mat[col][row] = mat[row][col]
		}
	}
	return mat
}

func Part1(input string) {
	startAt := "AA"
	valves := parseInput(input)
	solution := solvePart1(startAt, 30, valves)
	fmt.Println(solution)
}
