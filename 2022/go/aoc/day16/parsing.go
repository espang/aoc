package day16

import (
	"container/list"
	"strconv"
	"strings"
)

var testinput = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

type Valve struct {
	ID   string
	To   []string
	Rate int
}

func parseInput(input string) []Valve {
	lines := strings.Split(input, "\n")
	valves := []Valve{}
	for _, line := range lines {
		valves = append(valves, parseLine(line))
	}
	return valves
}

// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
func parseLine(line string) Valve {
	splitted := strings.Split(line, "; ")
	v, err := strconv.Atoi(splitted[0][23:])
	if err != nil {
		panic(err)
	}
	svs := strings.TrimPrefix(
		strings.TrimPrefix(splitted[1], "tunnels lead to valves "),
		"tunnel leads to valve ",
	)
	return Valve{
		ID:   line[6:8],
		To:   strings.Split(svs, ", "),
		Rate: v,
	}
}

func nonNullValves(valves []Valve) ([]string, []int) {
	names := []string{}
	rates := []int{}
	for _, k := range valves {
		if k.Rate > 0 {
			names = append(names, k.ID)
			rates = append(rates, k.Rate)
		}
	}
	return names, rates
}

func ShortestPath(from, target string, valves []Valve) int {
	distance := map[string]int{
		from: 0,
	}
	queue := list.New()
	queue.PushBack(from)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(string)
		dist := distance[current]
		var valve Valve
		for _, v := range valves {
			if v.ID == current {
				valve = v
			}
		}
		for _, to := range valve.To {
			if to == target {
				return dist + 1
			}
			if _, ok := distance[to]; ok {
				continue
			}
			distance[to] = dist + 1
			queue.PushBack(to)
		}
	}
	panic("not connected")
}

// NodeSet is an optimization that wasn't necessary;
// an attempt to improve the runtime before I found an issue
type NodeSet int16

func (ns NodeSet) Add(i int) NodeSet {
	return (ns | 1<<i)
}
func (ns NodeSet) Contains(i int) bool {
	return (ns & (1 << i)) > 0
}
