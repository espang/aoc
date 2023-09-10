package day11

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func (m *Monkey) ReceiveItem(item *big.Int) {
	m.items = append(m.items, item)
}

func step(monkeys []*Monkey) {
	three := big.NewInt(3)
	dummy := &big.Int{}
	zero := big.NewInt(0)
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			monkey.itemsInspected++
			worryLevel := monkey.operation(item)
			worryLevel = worryLevel.Div(worryLevel, three)
			dummy = dummy.Mod(worryLevel, monkey.divisbleBy)
			if dummy.Cmp(zero) == 0 {
				monkeys[monkey.ifTrue].ReceiveItem(worryLevel)
			} else {
				monkeys[monkey.ifFalse].ReceiveItem(worryLevel)
			}
		}
		monkey.items = monkey.items[0:0]
	}
}

func Part1(input string) {
	input = testInput
	monkeyStrings := strings.Split(input, "\n\n")
	var monkeys []*Monkey
	for _, ms := range monkeyStrings {
		monkeys = append(monkeys, parseMonkey(ms))
	}
	for i := 0; i < 20; i++ {
		step(monkeys)
	}
	var itemsInspected []int
	for _, monkey := range monkeys {
		itemsInspected = append(itemsInspected, monkey.itemsInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspected)))
	fmt.Println(itemsInspected[0] * itemsInspected[1])
}
