package day13

import (
	"strconv"
	"strings"
	"unicode"
)

func parseInteger(s string) (int, string) {
	end := 0
	for i, c := range s {
		if unicode.IsDigit(c) {
			end = i
		} else {
			break
		}
	}
	v, err := strconv.Atoi(s[0 : end+1])
	if err != nil {
		panic(err.Error())
	}
	return v, s[end+1:]
}

func endIndexOfList(s string) int {
	if s == "" {
		panic("endIndexOfList: empty string")
	}
	if s[0] != '[' {
		panic("endIndexOfList: has to start with [")
	}
	depth := 0
	for i, c := range s {
		switch c {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	panic("endIndexOfList: incomplete list: " + s)
}

func parseList(s string) ([]interface{}, string) {
	end := endIndexOfList(s)
	elements := parseAllElements(s[1:end])
	return elements, s[end+1:]
}

func parseNextElement(s string) (interface{}, string) {
	if s == "" {
		return nil, ""
	}
	switch s[0] {
	case '[':
		return parseList(s)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return parseInteger(s)
	case ',':
		return parseNextElement(s[1:])
	default:
		panic("unexpected string: " + s)
	}
}

func parseAllElements(s string) []interface{} {
	elements := []interface{}{}
	element, rest := parseNextElement(s)
	for element != nil {
		elements = append(elements, element)
		element, rest = parseNextElement(rest)
	}
	return elements
}

type Pair struct {
	list1 interface{}
	list2 interface{}
}

func parse(s string) []Pair {
	rawPairs := strings.Split(s, "\n\n")
	pairs := []Pair{}
	for _, p := range rawPairs {
		splitted := strings.Split(p, "\n")
		pairs = append(pairs, Pair{
			list1: parseAllElements(splitted[0]),
			list2: parseAllElements(splitted[1]),
		})
	}
	return pairs
}
