package day11

import (
	"math/big"
	"strconv"
	"strings"
)

var testInput = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

type Monkey struct {
	name           int
	items          []*big.Int
	operation      func(*big.Int) *big.Int
	divisbleBy     *big.Int
	ifTrue         int
	ifFalse        int
	itemsInspected int
}

func parseMonkey(s string) *Monkey {
	errorMsg := "unexpected monkey: " + s
	lines := strings.Split(s, "\n")
	if len(lines) != 6 {
		panic(errorMsg)
	}
	return &Monkey{
		name:       mustParseInt(mustTrimPreSuffix(lines[0], "Monkey ", ":", errorMsg), errorMsg),
		items:      mustParseInts(mustTrimPrefix(lines[1], "  Starting items: ", errorMsg), errorMsg),
		operation:  parseOperations(lines[2]),
		divisbleBy: mustParseBigInt(mustTrimPrefix(lines[3], "  Test: divisible by ", errorMsg), errorMsg),
		ifTrue:     mustParseInt(mustTrimPrefix(lines[4], "    If true: throw to monkey ", errorMsg), errorMsg),
		ifFalse:    mustParseInt(mustTrimPrefix(lines[5], "    If false: throw to monkey ", errorMsg), errorMsg),
	}
}

func mustTrimPrefix(s, prefix, msg string) string {
	if strings.HasPrefix(s, prefix) {
		return strings.TrimPrefix(s, prefix)
	}
	panic(msg)
}
func mustTrimSuffix(s, suffix, msg string) string {
	if strings.HasSuffix(s, suffix) {
		return strings.TrimSuffix(s, suffix)
	}
	panic(msg)
}
func mustTrimPreSuffix(s, prefix, suffix, msg string) string {
	return mustTrimPrefix(mustTrimSuffix(s, suffix, msg), prefix, msg)
}

func mustParseInt(val string, msg string) int {
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(msg)
	}
	return v
}
func mustParseBigInt(val string, msg string) *big.Int {
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(msg)
	}
	return big.NewInt(int64(v))
}
func mustParseInts(val string, msg string) []*big.Int {
	var ret []*big.Int
	for _, sv := range strings.Split(val, ", ") {
		ret = append(ret, mustParseBigInt(sv, msg))
	}
	return ret
}

func plus(v1, v2 *big.Int) *big.Int {
	return v1.Add(v1, v2)
}

func mul(v1, v2 *big.Int) *big.Int {
	return (&big.Int{}).Mul(v1, v2)
}

func parseOperations(s string) func(*big.Int) *big.Int {
	prefix := "  Operation: new = old "
	if !strings.HasPrefix(s, prefix) {
		panic("unexpected operations line: " + s)
	}
	splitted := strings.Split(strings.TrimPrefix(s, prefix), " ")
	if len(splitted) != 2 {
		panic("unexpected operations line: " + s)
	}

	var op func(*big.Int, *big.Int) *big.Int
	if splitted[0] == "+" {
		op = plus
	} else if splitted[0] == "*" {
		op = mul
	} else {
		panic("unexpected operations line: " + s)
	}

	if splitted[1] != "old" {
		val := mustParseBigInt(splitted[1], "unexpected operations line: "+s)
		return func(i *big.Int) *big.Int { return op(i, val) }
	} else if splitted[1] == "old" {
		return func(i *big.Int) *big.Int { return op(i, i) }
	} else {
		panic("unexpected operations line: " + s)
	}
}
