package day20

import (
	"fmt"
)

func Mix(nodes []*Node) {
	for _, n := range nodes {
		// optimization for part2:
		times := n.value % (len(nodes) - 1)
		for i := 0; i < times; i++ {
			SwapForwad(n)
		}
		for i := 0; i > times; i-- {
			SwapBackward(n)
		}
	}
}

func SumOfGroove(ring *Node) int {
	total := 0
	current := FindNodeByValue(ring, 0)
	for i := 1; i <= 3000; i++ {
		current = current.next
		if i%1000 == 0 {
			total += current.value
		}
	}
	return total
}

func Part1(input string) {
	numbers := Parse(input)
	ring := MakeRing(numbers)
	nodes := AllNodes(ring)
	Mix(nodes)
	fmt.Println(SumOfGroove(ring))
}
