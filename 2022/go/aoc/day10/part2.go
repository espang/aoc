package day10

import (
	"fmt"
	"strings"
)

type CRT struct {
	row    int
	values []rune
}

func (c CRT) String() string {
	ret := []string{}
	for i := 0; i < 6; i++ {
		ret = append(ret, string(c.values[i*40:(i+1)*40]))
	}
	return strings.Join(ret, "\n")
}

func (c *CRT) reset() {
	c.values = make([]rune, 240)
	for i := range c.values {
		c.values[i] = '.'
	}
}

func inSpirit(cycle, val int) bool {
	// val 1: ### ... (0, 1, 2)
	diff := (cycle % 40) - val
	return -1 <= diff && diff <= 1
}

func Part2(input string) {
	reg := registerFromInput(input)
	crt := &CRT{}
	crt.reset()
	for cycle, value := range reg.history {
		if inSpirit(cycle, value) {
			crt.values[cycle] = '#'
		}
	}

	fmt.Println()
	fmt.Println(crt)
}
