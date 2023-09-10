package day20

import (
	"fmt"

	"github.com/espang/aoc/y2022/go/aoc"
)

func Part2(input string) {
	numbers := aoc.Map(func(v int) int { return v * 811589153 }, Parse(input))
	ring := MakeRing(numbers)
	nodes := AllNodes(ring)
	for i := 0; i < 10; i++ {
		Mix(nodes)
	}
	fmt.Println(SumOfGroove(ring))
}
