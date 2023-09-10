package day20

import (
	"fmt"
)

func Mix(nodes []*Node) {
	for _, n := range nodes {
		movement := n.value % (len(nodes) - 1)
		positionToMoveTo := Move(movement, n)
		MoveAfter(n, positionToMoveTo)
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
	// PrintAll(nodes[0])
	fmt.Println(SumOfGroove(ring))
}
