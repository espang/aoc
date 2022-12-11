package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Register struct {
	value   int
	history []int
}

func (r *Register) handleAdd(x int) {
	r.history = append(r.history, r.value, r.value)
	r.value += x
}

func (r *Register) handleNoop() {
	r.history = append(r.history, r.value)
}

func registerFromInput(input string) *Register {
	reg := &Register{value: 1}
	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			reg.handleNoop()
		} else if strings.HasPrefix(line, "addx ") {
			v, err := strconv.Atoi(strings.TrimPrefix(line, "addx "))
			if err != nil {
				panic(err.Error())
			}
			reg.handleAdd(v)
		} else {
			panic("line: " + line)
		}
	}
	return reg
}

func Part1(input string) {
	reg := registerFromInput(input)
	total := 0
	for i := 20; i <= 220; i += 40 {
		total += i * reg.history[i-1]
	}
	fmt.Println(total)
}
