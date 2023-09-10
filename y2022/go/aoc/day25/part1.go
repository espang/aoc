package day25

import (
	"fmt"

	"github.com/espang/aoc/y2022/go/aoc"
)

func pow(base, exponent int) int {
	total := 1
	for i := 0; i < exponent; i++ {
		total *= base
	}
	return total
}

func toNumber(r rune) int {
	switch r {
	case '=':
		return -2
	case '-':
		return -1
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	}
	panic("unexpected rune")
}

func FromSnafu(s string) int {
	runes := []rune(s)
	total := 0
	for i := 0; i < len(runes); i++ {
		total += toNumber(runes[len(runes)-1-i]) * pow(5, i)
	}
	return total
}

func within(value int, middle, minMax int) bool {
	return value >= middle-minMax && value <= middle+minMax
}

func ToSnafu(i int) string {
	digits := []int{}
	bounds := []int{}
	max := 0
	for exponenent := 0; ; exponenent++ {
		val := pow(5, exponenent)
		digits = append(digits, val)
		bounds = append(bounds, max)
		max += 2 * val
		if 2*val > i {
			break
		}
	}
	result := ""
	rest := i
	for i := range digits {
		digit := digits[len(digits)-1-i]
		bound := bounds[len(digits)-1-i]

		if within(rest, 2*digit, bound) {
			result += "2"
			rest -= 2 * digit
		} else if within(rest, digit, bound) {
			result += "1"
			rest -= digit
		} else if within(rest, 0, bound) {
			result += "0"
		} else if within(rest, -digit, bound) {
			result += "-"
			rest += digit
		} else if within(rest, -2*digit, bound) {
			result += "="
			rest += 2 * digit
		} else {
			panic("")
		}
	}
	return result
}

func Part1(input string) {
	numbers := Parse(input)
	number := aoc.Sum(aoc.Map(FromSnafu, numbers))
	fmt.Println(ToSnafu(number))

}
func Part2(input string) {}
