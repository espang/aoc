package day16

import (
	"fmt"
)

func solvePart2(start string, maxTime int, valves []Valve) int {
	nodes, rates := nonNullValves(valves)
	distancesMatrix := distances(nodes, valves)
	timeToReachIFromStart := map[int]int{}
	for i, node := range nodes {
		timeToReachIFromStart[i] = 1 + ShortestPath(start, node, valves)
	}

	var solve func(int, int, int, int, NodeSet) int
	solve = func(p1Position, p1TimeLeft, p2Position, p2TimeLeft int, opened NodeSet) int {
		if p1TimeLeft < p2TimeLeft {
			p1Position, p2Position = p2Position, p1Position
			p1TimeLeft, p2TimeLeft = p2TimeLeft, p1TimeLeft
		}

		var max int
		for i := range nodes {
			if opened.Contains(i) {
				continue
			}
			timeLeft := p1TimeLeft - distancesMatrix[p1Position][i]
			if timeLeft > 0 {
				val := timeLeft*rates[i] + solve(i, timeLeft, p2Position, p2TimeLeft, opened.Add(i))
				if val > max {
					max = val
				}
			}
		}
		return max
	}
	var max int
	var opened NodeSet
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			p1TimeLeft := maxTime - timeToReachIFromStart[i]
			p2TimeLeft := maxTime - timeToReachIFromStart[j]
			val := p1TimeLeft*rates[i] +
				p2TimeLeft*rates[j] +
				solve(i, p1TimeLeft, j, p2TimeLeft, opened.Add(i).Add(j))
			if val > max {
				max = val
			}
		}
	}
	return max
}

func Part2(input string) {
	startAt := "AA"
	valves := parseInput(input)
	solution := solvePart2(startAt, 26, valves)
	fmt.Println(solution)
}
