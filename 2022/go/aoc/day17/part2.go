package day17

import "fmt"

func dropNRocks(n int, c *Chamber) {
	for i := 0; i < n; i++ {
		c.DropNextRock()
	}
}

// Part2 is hacky. The code to manually detect the cycle
// and its length is at the bottom.
// The code is run to find a pattern in the rocks. The code
// still needs 2 minutes to run.
func Part2(input string) {
	chamber := NewChamber(Parse(input))
	// number of rocks times number of streams
	patternLen := 5 * len(chamber.JetStream.pattern)

	// experimentation:
	// testinput:
	// 1 (7 pattern)
	// 308 (300 306 303 303 301 306 301)
	// 1 - 347 pattern for input
	times := (1_000_000_000_000 - patternLen) / (347 * patternLen)
	part := (1_000_000_000_000 - patternLen) % (347 * patternLen)

	// the socket until a repeating pattern forms
	dropNRocks(patternLen, &chamber)
	before := chamber.Highest
	dropNRocks(347*patternLen, &chamber)
	heightByPattern := chamber.Highest - before
	dropNRocks(part, &chamber)
	totalHeight := (times-1)*heightByPattern + chamber.Highest
	fmt.Println(totalHeight)

	// does it repeat every 5 * 10091 times?
	// after 347 applications of the 5*10091 drops it repeats the same pattern
	// values := []int{}
	// before := 0
	// for i := 0; i < 400; i++ {
	// 	for step := 0; step < patternLen; step++ {
	// 		chamber.DropNextRock()
	// 	}
	// 	values = append(values, chamber.Highest-before)
	// 	before = chamber.Highest
	// }
	// fmt.Println(values)
	// second := []int{}
	// first := values[0]
	// for _, v := range values[1:] {
	// 	second = append(second, first-v)
	// 	first = v
	// }
	// fmt.Println(second)
}
