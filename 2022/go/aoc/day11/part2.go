package day11

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func step2(monkeys []*Monkey, modulo *big.Int) {
	dummy := &big.Int{}
	zero := big.NewInt(0)
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			monkey.itemsInspected++
			worryLevel := monkey.operation(item)
			worryLevel = worryLevel.Mod(worryLevel, modulo)
			dummy.Mod(worryLevel, monkey.divisbleBy)
			if dummy.Cmp(zero) == 0 {
				monkeys[monkey.ifTrue].ReceiveItem(worryLevel)
			} else {
				monkeys[monkey.ifFalse].ReceiveItem(worryLevel)
			}
		}
		monkey.items = monkey.items[0:0]
	}
}

func Part2(input string) {
	monkeyStrings := strings.Split(input, "\n\n")
	var monkeys []*Monkey
	for _, ms := range monkeyStrings {
		monkeys = append(monkeys, parseMonkey(ms))
	}

	theModulo := big.NewInt(1)
	for _, monkey := range monkeys {
		theModulo = theModulo.Mul(theModulo, monkey.divisbleBy)
	}

	for i := 0; i < 10000; i++ {
		step2(monkeys, theModulo)
	}
	var itemsInspected []int
	for _, monkey := range monkeys {
		itemsInspected = append(itemsInspected, monkey.itemsInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspected)))
	fmt.Println(itemsInspected[0] * itemsInspected[1])
}
