package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/espang/aoc/aoc"
)

type Rule struct {
	Page1, Page2 int
}

func RuleFromLine(s string) Rule {
	parts := strings.Split(s, "|")
	if len(parts) != 2 {
		panic("RuleFromLine")
	}
	page1, _ := strconv.Atoi(parts[0])
	page2, _ := strconv.Atoi(parts[1])
	return Rule{Page1: page1, Page2: page2}
}

func Check(pages []int) func(Rule) bool {
	return func(r Rule) bool {
		rulePages := aoc.Filter(func(p int) bool {
			return p == r.Page1 || p == r.Page2
		}, pages)
		if len(rulePages) == 2 {
			return rulePages[0] == r.Page1
		}
		return true
	}
}

type Update []int

func UpdateFromLine(s string) Update {
	return aoc.Make(
		func(s string) int {
			v, _ := strconv.Atoi(s)
			return v
		},
		aoc.SplitByComma(s),
	)
}

func (u Update) Check(rules []Rule) bool {
	return aoc.All(Check(u), rules)
}

func (u Update) Append(page int, rules []Rule) Update {
	for i, p := range u {
		for _, r := range rules {
			if r.Page1 == page && r.Page2 == p {
				return append(u[:i], append([]int{page}, u[i:]...)...)
			}
		}
	}
	return append(u, page)
}

func parse(s string) ([]Rule, []Update) {
	parts := strings.Split(s, "\n\n")
	if len(parts) != 2 {
		panic("")
	}

	rules := aoc.Make(RuleFromLine, aoc.SplitByLine(parts[0]))
	updates := aoc.Make(UpdateFromLine, aoc.SplitByLine(parts[1]))
	return rules, updates
}

func Part1(s string) {
	rules, updates := parse(s)
	total := 0
	for _, u := range updates {
		if u.Check(rules) {
			total += u[len(u)/2]
		}
	}
	fmt.Println(total)
}

func (u Update) Fix(rs []Rule) Update {
	var incorrects []int
	var corrects Update
outer:
	for i, page1 := range u {
	inner:
		for _, page2 := range u[i+1:] {
			for _, r := range rs {
				if r.Page1 == page1 && r.Page2 == page2 {
					continue inner
				}
				if r.Page1 == page2 && r.Page2 == page1 {
					incorrects = append(incorrects, page1)
					continue outer
				}
			}
		}
		corrects = append(corrects, page1)
	}
	for _, page := range incorrects {
		corrects = corrects.Append(page, rs)
	}
	return corrects
}

func Part2(s string) {
	rules, updates := parse(s)
	total := 0
	for _, u := range updates {
		if !u.Check(rules) {
			u = u.Fix(rules)
			total += u[len(u)/2]
		}
	}
	fmt.Println(total)
}
