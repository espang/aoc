package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/espang/aoc/aoc"
)

type Equation struct {
	Result  int
	Numbers []int
}

type Operation func(int, int) int

func Plus(lhs, rhs int) int {
	return lhs + rhs
}

func Mul(lhs, rhs int) int {
	return lhs * rhs
}

func Concat(lhs, rhs int) int {
	v, _ := strconv.Atoi(strconv.Itoa(lhs) + strconv.Itoa(rhs))
	return v
}

func (eq Equation) Possible(operations []Operation) bool {
	values := []int{eq.Numbers[0]}
	for _, n := range eq.Numbers[1:] {
		var newValues []int
		for _, op := range operations {
			for _, val := range values {
				newValues = append(newValues, op(val, n))
			}
		}
		values = newValues
	}
	return slices.Contains(values, eq.Result)
}

func ParseLine(s string) Equation {
	parts := strings.Split(s, ":")
	result, _ := strconv.Atoi(parts[0])
	numbers := aoc.Map(func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}, strings.Split(strings.TrimSpace(parts[1]), " "))
	return Equation{
		Result:  result,
		Numbers: numbers,
	}
}

func Parse(s string) []Equation {
	return aoc.Map(ParseLine, aoc.SplitByLine(s))
}

func Part1(s string) {
	equations := Parse(s)
	total := 0
	operations := []Operation{Plus, Mul}
	for _, eq := range equations {
		if eq.Possible(operations) {
			total += eq.Result
		}
	}
	fmt.Println(total)
}

func Part2(s string) {
	equations := Parse(s)
	total := 0
	operations := []Operation{Plus, Mul, Concat}
	for _, eq := range equations {
		if eq.Possible(operations) {
			total += eq.Result
		}
	}
	fmt.Println(total)
}
